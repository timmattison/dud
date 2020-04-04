package cache

import (
	"fmt"
	"github.com/kevlar1818/duc/artifact"
	"github.com/kevlar1818/duc/fsutil"
	"github.com/kevlar1818/duc/strategy"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
)

// Commit calculates the checksum of the artifact, moves it to the cache, then performs a checkout.
func (cache *LocalCache) Commit(workingDir string, art *artifact.Artifact, strat strategy.CheckoutStrategy) error {
	srcPath := path.Join(workingDir, art.Path)
	isRegFile, err := fsutil.IsRegularFile(srcPath)
	if err != nil {
		return err
	}
	if !isRegFile {
		return fmt.Errorf("file %#v is not a regular file", srcPath)
	}
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err // Don't wrap this so we can use os.IsNotExist on it
	}
	defer srcFile.Close()
	dstFile, err := ioutil.TempFile(cache.Dir, "")
	if err != nil {
		return errors.Wrapf(err, "creating tempfile in %#v failed", cache.Dir)
	}
	defer dstFile.Close()

	// TODO: only copy if the cache is on a different filesystem (os.Rename if possible)
	// OR, if we're using CopyStrategy
	checksum, err := fsutil.ChecksumAndCopy(srcFile, dstFile)
	if err != nil {
		return errors.Wrapf(err, "checksum of %#v failed", srcPath)
	}
	dstDir := path.Join(cache.Dir, checksum[:2])
	if err = os.MkdirAll(dstDir, 0755); err != nil {
		return errors.Wrapf(err, "mkdirs %#v failed", dstDir)
	}
	cachePath := path.Join(dstDir, checksum[2:])
	if err = os.Rename(dstFile.Name(), cachePath); err != nil {
		return errors.Wrapf(err, "mv %#v failed", dstFile)
	}
	if err := os.Chmod(cachePath, 0444); err != nil {
		return errors.Wrapf(err, "chmod %#v failed", cachePath)
	}
	art.Checksum = checksum
	// There's no need to call Checkout if using CopyStrategy; the original file still exists.
	if strat == strategy.LinkStrategy {
		// TODO: add rm to checkout as "force" option
		if err := os.Remove(srcPath); err != nil {
			return errors.Wrapf(err, "rm %#v failed", srcPath)
		}
		return cache.Checkout(workingDir, art, strat)
	}
	return nil
}

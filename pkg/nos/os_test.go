package nos

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

func TestCopy(t *testing.T) {
	{
		// test/temp/pkg does not exist
		// so Copy nos to pkg will be a clone
		cleanTmpDir()
		src := "."
		dst := path.Join(tmpDir, "pkg")

		Copy(src, dst)
		srcPaths, err := AbsPaths(src)
		assert.Nil(t, err)
		dstPaths, err := AbsPaths(dst)
		assert.Nil(t, err)
		for i := range dstPaths {
			srcPaths[i] = path.Base(srcPaths[i])
			dstPaths[i] = path.Base(dstPaths[i])
		}
		assert.Equal(t, "nos", srcPaths[0])
		assert.Equal(t, "pkg", dstPaths[0])
		assert.Equal(t, srcPaths[1:], dstPaths[1:])
	}
	{
		// test/temp/pkg does exist
		// so Copy nos to pkg will be an into op
		cleanTmpDir()
		src := "."
		dst := path.Join(tmpDir, "pkg")
		MkdirP(dst)

		Copy(src, dst)
		srcPaths, err := AbsPaths(src)
		assert.Nil(t, err)
		dstPaths, err := AbsPaths(path.Join(dst, "nos"))
		assert.Nil(t, err)
		for i := range dstPaths {
			srcPaths[i] = path.Base(srcPaths[i])
			dstPaths[i] = path.Base(dstPaths[i])
		}
		assert.Equal(t, srcPaths, dstPaths)
	}
}

func TestCopyFile(t *testing.T) {
	{
		cleanTmpDir()
		foo := path.Join(tmpDir, "foo")

		assert.False(t, Exists(foo))
		CopyFile(readme, foo)
		assert.True(t, Exists(foo))

		srcMD5, _ := MD5(readme)
		dstMD5, _ := MD5(foo)
		assert.Equal(t, srcMD5, dstMD5)
	}
}

func TestExists(t *testing.T) {
	assert.False(t, Exists("bob"))
	assert.True(t, Exists(readme))
}

func TestIsDir(t *testing.T) {
	assert.False(t, IsDir(readme))
	assert.True(t, IsDir("../.."))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile(readme))
	assert.False(t, IsFile("../.."))
}

func TestMD5(t *testing.T) {
	if Exists(tmpfile) {
		os.Remove(tmpfile)
	}
	f, _ := os.Create(tmpfile)
	defer f.Close()
	f.WriteString(`This is a test of the emergency broadcast system.`)

	expected := "067a8c38325b12159844261d16e5cb13"
	result, _ := MD5(tmpfile)
	assert.Equal(t, expected, result)
}

func TestMkdirP(t *testing.T) {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
	assert.True(t, Exists(tmpDir))
}

func TestReadLines(t *testing.T) {
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
}

func TestTouch(t *testing.T) {
	cleanTmpDir()
	assert.False(t, Exists(tmpfile))
	assert.Nil(t, Touch(tmpfile))
	assert.True(t, Exists(tmpfile))
	assert.Nil(t, Touch(tmpfile))
}

func TestWriteLines(t *testing.T) {
	cleanTmpDir()
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
	err = WriteLines(tmpfile, lines)
	assert.Nil(t, err)
	{
		lines2, err := ReadLines(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, lines, lines2)
	}
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}

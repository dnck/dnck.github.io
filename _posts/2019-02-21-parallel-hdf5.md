---
layout: post
title: Parallel HDF5
author: dc
date: 2019-02-21
comments: true
analytics: true
keywords:  HDF5
description: installation of parallel hdf5 on an ubuntu server
category: blog
show_excerpt: true
excerpt: Failed step-by-step instructions to install parallel hdf5
---

This is the second time in less than a year I've face the problem of configuring parallel hdf5 on a server. I don't even know why I bother! Oh, yeah, that's right, I thought hdf5 was cool when I was first introduced to it, and I wanted to make it even more complicated by introducing concurrent i/o. Anyway, what follows is my catalogue of errors when trying to install on an ubuntu server. If you've got any suggestions on what I'm doing wrong, please comment :().

### What do?

Attempting to .configure, make, make install, OpenMpi & parallel-HDf5 (using the CC=mpicc option)
Attempting to .configure, make, make install zlib
Attempting to pip install mpi4py & h5py (with links to the pHDf5 library using specific options for pip)

See below for details on steps already taken and specific bugs:

### OPENMPI
Status: No bugs
```
> curl -O https://download.open-mpi.org/release/open-mpi/v2.0/openmpi-2.0.4.tar.gz
> tar xf openmpi-2.0.4.tar.gz
> cd openmpi-2.0.4
> ./configure --prefix=$HOME/.local
> make all
> make install
> mpirun --version
```
This compiles perfectly. All the executables are located in my .local path.

The last commands prints,
```
> mpirun (Open MPI) 2.0.4
```
### ZLIB (shared object)
Status: Bugs
```
> curl -O ...
> tar xf ...
> cd zlib-x.x.x
> ./configure --prefix=$HOME/.local
> make test
```

Everything works. The C header file is in the .local/include dir and the shared object is in the .local/lib directory.

### ZLIB (Python)
Status: bugs

However, installation of the zlib Python module with pip fails:
```
> pip install zlib
```
This prints,
```
> Collecting zlib
>    Could not find a version that satisfies the requirement zlib (from versions: )
>    No matching distribution found for zlib
```
I tried to fix this by explicitly telling the C compiler where to find the zlib stuff.
```
> zlib_lib="$HOME/.local/lib"
> zlib_inc="$HOME/.local/include"
> export CPPFLAGS="-I${zlib_inc} ${CPPFLAGS}"
> export LD_LIBRARY_PATH="${zlib_lib}:${LD_LIBRARY_PATH}"
> export LDFLAGS="-L${zlib_lib} -Wl,-rpath=${zlib_lib} ${LDFLAGS}"
> pip install zlib
```
But, this also fails and produces the same message as above from pip.

So, it seems I have zlib properly installed, but I do not have a working Python module called, zlib.

Oh, I'm running the default ubuntu Python 3.6 (not anaconda or anything.)

### pHDf5 (C)
Status: bugs

Downloaded from https://www.hdfgroup.org/downloads/hdf5/source-code/
```
> cd hdf5-1.10.4
> CC=$HOME/.local/bin/mpicc ./configure --prefix=$HOME/.local --enable-shared --enable-parallel
```
This seems to give me the output I want. But doing
```
> make
> make check
```
fails to return a positive result. Namely, the MPI_ABORT seems to be called during the parallel tests.

Also, I have tried to install the regular hdf5 library (with no parallel support) by doing,
```
> CC=$HOME/.local/bin/mpicc ./configure --prefix=$HOME/.local --enable-shared
```
But, this also fails on the > make check stage.

Note that I have gotten this setup running before with my macbook pro 8.1 back in May 2018:
https://www.danjcook.com/blog/2018/05/16/parallel-HDF5-for-use-in-Python.html

The above link might be helpful in future attempts to get this constellations of software up and running on the master-node.

End pain and suffering, please help.



A common coin is a "distributed object that is expected to deliver, with some probability, the same sequence of random bits b1, b2, ...,bn ...to each process"

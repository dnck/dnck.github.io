---
layout: post
title: Parallel computing with HDF5, MPI and Python
date: 2018-05-16
author: dc
comments: true
analytics: false
keywords: parallel, HDF5, MPI, python
description: how I installed parallel HDF5 for use in Python with MPI
tag: parallel, HDF5, MPI
category: blog
---


The following post details how I got a parallel HDF5 setup with Python bindings running on my old macbook pro 8.1. 

TLDR: It's not easy, and the installation might be OS and machine specific. Anyway, here's the steps I took.

### Step 1. Install openmpi

I downloaded version 2.04 of openmpi from https://www.open-mpi.org/software/ompi/v2.0/

On terminal, 
```
tar xf openmpi-2.0.4.tar.gz
cd openmpi-2.0.4/
./configure --prefix=$HOME/openmpi/
make all
make install
$HOME/mpirun/mpirun --version
```
The last line prints,
```
mpirun (Open MPI) 2.0.4.
```

### Step 2. Install HDF5, the parallel version

Download HDF5 from their website, https://www.hdfgroup.org/downloads/hdf5/source-code/

On terminal,
```
CC=$HOME/mpicc ./configure --with-zlib=/usr/local/opt --disable-fortran --prefix=$HOME --enable-shared --enable-parallel
make
make check
sudo make install
h5pcc -showconfig
```
That ends up showing some configuration variables for HDF5.

Note that I'm not sure about the disable-fortan option. But, I don't use fortran, and I saw this suggsted somewhere on the internet. So I went with it. Note also that I had install zlib with homebrew. But, on Ubuntu, with some restricted permissions, I couldn't get this running.

### Step 3. Install mpi4py

I installed mpi4py with pip using the mpicc compiler option set:
```
env MPICC=$HOME/mpicc pip install mpi4py
```
On mac, I had trouble finding the right path to the mpicc. So, if you're using this, just make sure you got the compiler's path right.

### Step 4. Install h5py

I had installed h5py before, so I uninstalled my old version of h5py:
```
pip uninstall h5py
```
and then I reinstalled with pip using some options, 
```
CC="mpicc" HDF5_MPI="ON" HDF5_DIR=/usr/local/bin/ pip install --no-binary=h5py h5py
```
Again, this one gave me some trouble. But, if you get the paths correct, it shouldn't be a problem. 

How did it work?

### Step 5. Test it in Python

Of course, I wanted to know whether everything was working in Python. So, I fired up my notebook, and ... 
```
from mpi4py import MPI
print("Hello World (from process %d)" % MPI.COMM_WORLD.Get_rank())
```
The last line prints, 
```
Hello World (from process 0)
```
I also wanted to make sure h5py was working, 
```
import h5py
print h5py.version.info
```
which prints a summary of the h5py configuration. 

Everything looks dandy.

My final test,
```
rank = MPI.COMM_WORLD.rank  # The process ID (integer 0-3 for 4-process run)
f = h5py.File('parallel_test.hdf5', 'w', driver='mpio', comm=MPI.COMM_WORLD)
dset = f.create_dataset('test', (4,), dtype='i')
dset[rank] = rank
f.close()
import os
print(os.listdir(os.getcwd()))
```
prints,
```
parallel_test.hdf5
```
which implies I successfully wrote an hdf5 file with this parallel setup. 

### Conclusion

Configuring HDF5 with parallel support from OpenMPI, and the Python modules, mpi4py and h5py for parallel processing is not very straight forward from someone not already well-versed in compiling C programs, and even with, it seems like there are a lot of gotchas. 


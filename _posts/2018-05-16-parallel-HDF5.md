---
layout: post
title: OpenMPI, parallel HDF5, mpi4py, and h5py
date: 2018-05-16
author: dc
comments: true
analytics: false
keywords: OpenMPI, parallel, HDF5, mpi4py, and h5py, pythong
description: how I installed parallel HDF5 for use in Python with MPI
tag: parallel, HDF5, MPI
category: blog
show_excerpt: true
excerpt: Step-by-step install instructions for parallel big-data processing 
---

Here's the steps I took to set up pHDF5, MPI and Python on my macOS.

**Openmpi**

```
curl -O https://download.open-mpi.org/release/open-mpi/v2.0/openmpi-2.0.4.tar.gz
tar xf openmpi-2.0.4.tar.gz
cd openmpi-2.0.4/
./configure --prefix=usr/local/openmpi
make all
make install
usr/local/openmpi/mpirun/mpirun --version
```

If successful, the last line prints,

```
mpirun (Open MPI) 2.0.4.
```

Notes
* Get the link for the openmpi download from https://www.open-mpi.org/
* This installs to the usr/local directory


**HDF5, the parallel version**

Download parallel HDF5 to ```usr/local/```, unzip, and cd

```
CC=$HOME/mpicc ./configure --with-zlib=/usr/local/opt --disable-fortran --prefix=$HOME --enable-shared --enable-parallel
make
make check
sudo make install
h5pcc -showconfig
```

That ends up showing you some configuration variables for HDF5.

If all goes well, you should see you got the parallel version.

**mpi4py**

I installed mpi4py via,

```
env MPICC=usr/local/openmpi/mpicc pip install mpi4py
```

**h5py**

I uninstalled my older version of h5py:

```
pip uninstall h5py
```

and I reinstalled via,

```
CC="mpicc" HDF5_MPI="ON" HDF5_DIR=/usr/local/bin/ pip install --no-binary=h5py h5py
```

**Test**

```
from mpi4py import MPI
print "Hello World (from process %d)" % MPI.COMM_WORLD.Get_rank()
```

The last line prints,

```
Hello World (from process 0)
```

```
import h5py
print h5py.version.info
```

The last line prints a summary of the h5py configuration. Everything looked dandy.

Also,

```
rank = MPI.COMM_WORLD.rank  # The process ID (integer 0-3 for 4-process run)
f = h5py.File('parallel_test.hdf5', 'w', driver='mpio', comm=MPI.COMM_WORLD)
dset = f.create_dataset('test', (4,), dtype='i')
dset[rank] = rank
f.close()
import os
print os.listdir(os.getcwd())
```

The last line prints,

```
parallel_test.hdf5
```

which implies you got a hdf5 file.

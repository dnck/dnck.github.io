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
category: parallel, HDF5, MPI
---

My previous post about Tensorflow gave me some ideas.

Let's do parallel computing the right way. Let's use HDF5 along with MPI for parallel processing.

Here's what I did to get the system up and running on my MAC OS.

I'm sure it is a hack job, but it works!

  Step 1. Install openmpi

I downloaded version 2.04 of openmpi from https://www.open-mpi.org/software/ompi/v2.0/

After, I went to terminal, and did,

```
tar xf openmpi-2.0.4.tar.gz
cd openmpi-2.0.4/
./configure --prefix=$HOME/openmpi/
make all
make install
$HOME/mpirun/mpirun --version
```

For some reason that created an opt/ folder in my $HOME path.

Whatever. It works. The last line printed,

```
mpirun (Open MPI) 2.0.4.
```

## Step 2. install HDF5, the parallel version

Next, I installed the HDF5 parallel version:

```
CC=$HOME/mpicc ./configure --with-zlib=/usr/local/opt --disable-fortran --prefix=$HOME --enable-shared --enable-parallel
make
make check
sudo make install
h5pcc -showconfig
```

That ends up showing you some configuration variables for HDF5.

If all goes well, you should see you got the parallel version.

## Step 3. install mpi4py

I installed mpi4py via,

```
env MPICC=$HOME/mpicc pip install mpi4py
```

That didn't give me much, but I did have some trouble passing the right path to the MPICC. Just make sure you feed it the compiler's path.

## Step 4. Install h5py

Last, I uninstalled my older version of h5py:

```
pip uninstall h5py
```

 and then i reinstalled via,

```
CC="mpicc" HDF5_MPI="ON" HDF5_DIR=/usr/local/bin/ pip install --no-binary=h5py h5py
```

Again, this one gave me some trouble. I found numerous variations of the CC=, HDF5_MPI=, HDF5_DIR= commands all over the place. Anyway, this is the one that worked for me.

How did it work?

## Step 5. Test it

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

**Thus, next comes the real test:**

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

which implies I got a hdf5 file!

## Conclusion

Thus, it looks like the first output works as planned!

I hope the same sort of logic will apply on the Ubuntu computer at Helix!!

My goal is to parallelize the shit of out some massive computations I'm about to be doing. Next step. Docker containers, then Azure, then 1 billion machines running our communities code to cast an anonymous shadow over the conglomerate data broker's flourishing gardens.

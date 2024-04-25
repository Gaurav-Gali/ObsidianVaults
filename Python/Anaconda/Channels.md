They are the entities that hold the packages required for installing

**Viewing all channels**

```shell
conda config --show channels
```

**Installing packages from different channels**

```shell
conda install -c pytorch pytorch
```

_Note_ : The first pytorch here is the channel name and the second one is the package.

**Adding channels to the list of channels**

```shell
conda config --add channels pytorch
```

_Note_ : The pytorch here is the channel not the package.

**An Important channel for wider range of packages**

```shell
conda config --add channels conda-forge
```

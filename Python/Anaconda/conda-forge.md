
1. Add the Conda Forge channel
```
conda config --add channels conda-forge
```

2. Set Conda Forge to be the priority channel
```
conda config --set channel_priority strict
```

3. Installing packages
```
conda install <package-name> -c conda-forge
```
==Note== : since conda-forge is set as the priority channel , we can directly use 
```
conda install <package-name>
```

4. Verifying installation
```
conda config --show channels
```

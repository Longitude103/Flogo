# Flogo

This is a go implementation that is inspired by the [Flopy](https://github.com/modflowpy/flopy) package
and is intended to help write out [MODFLOW 6](https://www.usgs.gov/software/modflow-6-usgs-modular-hydrologic-model) 
that can be found in this [repo](https://github.com/MODFLOW-USGS/modflow6) files for model execution. This package
is not intended to be run directly, but integrated into other GO programs that require MODFLOW
file outputs.

## Initial Release
The initial release will preform one activity which is to create output WEL6 and RCH files used by 
MODFLOW 6.

### File Creation
The initial release will create WEL and RCH files in MODFLOW 6 format. New releases
will seek to be able to support all the model files and download and execute the code
to run the model. 

The MODFLOW 6 files conform to follow the following document [MODFLOW 6 Description of Input/Output](https://water.usgs.gov/water-resources/software/MODFLOW-6/mf6io_6.2.2.pdf)

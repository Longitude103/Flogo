# Flogo

This is a go implementation that is inspired by the Flopy package in Python
and is intended to help write out MODFLOW 6 files for model execution. This package
is not intended to be run directly, but integrated into other GO programs that require MODFLOW
file outputs.

## Initial Release
The initial release will do two limited activities, create some initial output files used by 
MODFLOW and be able to analyse one of the output file.

### File Creation
The initial release will create WEL and RCH files in MODFLOW 6 format. New releases
will seek to be able to support all the model files and download and execute the code
to run the model. 

### Output file analysis
The initial release will create an analysis of the MODFLOW 6 results and create graphs, charts 
and Excel files of the results.

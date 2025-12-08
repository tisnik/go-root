OUTFILE="go.times"
PREFIX="mandelbrot"
RESOLUTION="4096"
MAXITER="255"

rm $OUTFILE

for i in `seq 10`
do
    echo $i
    /usr/bin/time --output $OUTFILE --append --format "%e %M" ./mandelbrot $RESOLUTION $RESOLUTION $MAXITER > "${PREFIX}_${i}.ppm"
done

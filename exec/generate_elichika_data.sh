# shell scripts to generate data for elichika, from various input sources, by the order of certainty about accuracy

go build -o converter ./exec && \
echo "doc input" >log.txt && \
./converter undefined doc input/doc elichika output json 2>>log.txt && \
echo "triangle input" >>log.txt && \
./converter undefined triangle input/triangle elichika output json 2>>log.txt && \
echo "Done!"
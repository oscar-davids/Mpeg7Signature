#!/bin/bash
for i in {1..10}
do
	./livepeer_bench  -in bbb/source.m3u8 -transcodingOptions transcodingOptions.json -concurrentSessions $i -nvidia 0 -outPrefix tmp1/ -sign 
done


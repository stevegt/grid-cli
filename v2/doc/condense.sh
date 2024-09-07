#!/bin/bash

set -x

# padsp signalgen -t 100m sin 444

qfile=/tmp/qfile.txt
pfile=/tmp/$$.prompt

while true
do
    chatfile=/tmp/$$.chat
    rm -f $chatfile

    grok aidda commit 
    # padsp signalgen -t 100m sin 600

    inputs=$(ls *.md)
    inputs=$(echo $inputs | sed 's/ /,/g')

    sysmsg="You are an expert technical writer and software architect. 
        We are writing a book about PromiseGrid.  Each of the input files
        is a chapter in the book."

    prompt="Examine the chapters.  Which two chapters are most
    identical and could b merged into one?  The first line of your
    answer should be the heading 'File1:' followed by the filename of
    the first chapter to be merged.  The second line of your answer
    should be the heading 'File2:' followed by the filename of the
    second chapter to be merged. The rest of the answer should be a
    description of changes to make to File2 to make it identical to
    File1."

    echo $prompt | grok chat $chatfile -s "$sysmsg" -i $inputs > $qfile

    fn1=$(head -n 1 $qfile | sed 's/File1: //')
    fn2=$(head -n 2 $qfile | tail -n 1 | sed 's/File2: //')

    if [ -z "$fn1" ] || [ -z "$fn2" ]; then
        echo "filenames not found"
        padsp signalgen -t 100m sin 300
        continue
    fi

    if [ "$fn1" == "$fn2" ]; then
        echo "filenames are the same"
        padsp signalgen -t 100m sin 300
        continue
    fi

    if [ ! -f $fn1 ] || [ ! -f $fn2 ]; then
        echo "files not found"
        padsp signalgen -t 100m sin 300
        continue
    fi

    echo "Modify $fn2 according to the following instructions:" > $pfile
    echo >> $pfile
    tail -n +3 $qfile >> $pfile

    grok chat $chatfile -s "$sysmsg" -i $fn1,$fn2 -o $fn2 < $pfile
    padsp signalgen -t 100m sin 500
    inotifywait -e modify $fn
done





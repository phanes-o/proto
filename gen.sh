#!/bin/bash
work_path=$(dirname "$PWD")

function read_dir(){
    for file in `ls -a $1`
    do
        if [ -d $1"/"$file ]
        then
            if [[ $file != '.' && $file != '..' && $file != '.git' ]]
            then
              read_dir $1"/"$file $1
            fi
        else
          if [[ "${file##*.}"x == 'proto'x ]]
          then
              protoc -I=$(dirname $2) --go_out=paths=source_relative:../ --micro_out=paths=source_relative:../ $1"/"$file
          fi
        fi
    done
}

read_dir $work_path $work_path

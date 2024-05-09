#!/bin/sh

args="Lorem ipsum dolor sit amet"

run() {
	echo "==== Running '$(pwd)/client1 ${args}' ===="
	./client1 ${args}
	echo
	echo "=========================="

	echo

	echo "==== Running '$(pwd)/client2 ${args}' ===="
	./client2 ${args}
	echo
	echo "=========================="
}

cd /unpatched
run
echo

cd /patched
run


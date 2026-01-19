# go archive example

This repository was copied from [vladimirvivien/go-cshared-examples](https://github.com/vladimirvivien/go-cshared-examples) and is meant to demonstrate the current issue with compiling go archives for non glibc linux distros.

The current issue to resolve this lives here [github.com/golang/go/pull/67254](https://github.com/golang/go/pull/67254)

## testing

This repo contains 2 c programs, client 1 is linked at compile time and client 2 is dynamically loading the go archive.

I have a patch to partially fix this issue [github.com/golang/go/pull/67254](https://github.com/golang/go/pull/67254) that is applied to the go source tree inside both docker containers. Both clients are first compiled with an unpatched go source tree and again with a patched source tree to ensure the ubuntu (glibc) behaves properly with the patch applied.


### ubuntu docker image

`make ubuntu` builds the shared archives for ubuntu and demonstrates the expected behavior where both clients are able to call go functions and go is able to read the program arguments.

example output:

```txt
==== Running '/unpatched/client1 Lorem ipsum dolor sit amet' ====
Using awesome lib from C:
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(77,12,5,99,28,23): 5,12,23,28,77,99,
Hello from C!
Program has 6 arguments
arg 0: ./client1
arg 1: Lorem
arg 2: ipsum
arg 3: dolor
arg 4: sit
arg 5: amet
Program has 8 environment variables
env 0: HOSTNAME=0782ec67a336
env 1: SHLVL=0
env 2: HOME=/root
env 3: OLDPWD=/
env 4: _=./entrypoint.sh
env 5: TERM=xterm
env 6: PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
env 7: PWD=/unpatched

==========================

==== Running '/unpatched/client2 Lorem ipsum dolor sit amet' ====
awesome.Add(12, 99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(44,23,7,66,2): 2,7,23,44,66,
Hello from C!
Program has 6 arguments
arg 0: ./client2
arg 1: Lorem
arg 2: ipsum
arg 3: dolor
arg 4: sit
arg 5: amet
Program has 8 environment variables
env 0: HOSTNAME=0782ec67a336
env 1: SHLVL=0
env 2: HOME=/root
env 3: OLDPWD=/
env 4: _=./entrypoint.sh
env 5: TERM=xterm
env 6: PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
env 7: PWD=/unpatched

==========================

==== Running '/patched/client1 Lorem ipsum dolor sit amet' ====
Using awesome lib from C:
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(77,12,5,99,28,23): 5,12,23,28,77,99,
Hello from C!
Program has 6 arguments
arg 0: ./client1
arg 1: Lorem
arg 2: ipsum
arg 3: dolor
arg 4: sit
arg 5: amet
Program has 8 environment variables
env 0: HOSTNAME=0782ec67a336
env 1: SHLVL=0
env 2: HOME=/root
env 3: OLDPWD=/unpatched
env 4: _=./entrypoint.sh
env 5: TERM=xterm
env 6: PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
env 7: PWD=/patched

==========================

==== Running '/patched/client2 Lorem ipsum dolor sit amet' ====
awesome.Add(12, 99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(44,23,7,66,2): 2,7,23,44,66,
Hello from C!
Program has 6 arguments
arg 0: ./client2
arg 1: Lorem
arg 2: ipsum
arg 3: dolor
arg 4: sit
arg 5: amet
Program has 8 environment variables
env 0: HOSTNAME=0782ec67a336
env 1: SHLVL=0
env 2: HOME=/root
env 3: OLDPWD=/unpatched
env 4: _=./entrypoint.sh
env 5: TERM=xterm
env 6: PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
env 7: PWD=/patched

==========================
```

### alpine docker image

`make alpine` builds the shared archives for alpine and demonstrates the issue before and after the patch has been applied

example output:

```txt
==== Running '/unpatched/client1 Lorem ipsum dolor sit amet' ====
Using awesome lib from C:
Segmentation fault (core dumped)

==========================

==== Running '/unpatched/client2 Lorem ipsum dolor sit amet' ====
Error relocating ./awesome.so: free: initial-exec TLS resolves to dynamic definition in ./awesome.so
==========================

==== Running '/patched/client1 Lorem ipsum dolor sit amet' ====
Using awesome lib from C:
awesome.Add(12,99) = 111
awesome.Cosine(1) = 0.540302
awesome.Sort(77,12,5,99,28,23): 5,12,23,28,77,99,
Hello from C!
Program has 6 arguments
arg 0: ./client1
arg 1: Lorem
arg 2: ipsum
arg 3: dolor
arg 4: sit
arg 5: amet
Program has 7 environment variables
env 0: HOSTNAME=9d24b58d8bd1
env 1: SHLVL=2
env 2: HOME=/root
env 3: OLDPWD=/unpatched
env 4: TERM=xterm
env 5: PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
env 6: PWD=/patched

==========================

==== Running '/patched/client2 Lorem ipsum dolor sit amet' ====
Error relocating ./awesome.so: free: initial-exec TLS resolves to dynamic definition in ./awesome.so
==========================
```
# peerid-locker

[![Build Status](https://travis-ci.com/RTradeLtd/peerid-locker.svg?branch=master)](https://travis-ci.com/RTradeLtd/peerid-locker) [![codecov](https://codecov.io/gh/RTradeLtd/peerid-locker/branch/master/graph/badge.svg)](https://codecov.io/gh/RTradeLtd/peerid-locker) [![GoDoc](https://godoc.org/github.com/RTradeLtd/peerid-locker?status.svg)](https://godoc.org/github.com/RTradeLtd/peerid-locker)

peerid-locker is used to manage concurrent access to resources that are peerID specific, such as when publishing IPNS records. It enables us to lock access to a resource that is specific to a particular peerID, without blocking access to resources that are specific to other peerIDs
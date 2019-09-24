/*
Package peeridlocker is used to provide granular concurrent resource access control on a per-peerID.

It is particularily useful in things like the IPFS namesys package,  which completely blocks access to
record publishing, any time a record is being published. This is done to ensure that we can atomically increment
sequence number, however it has the adverse effect of blocking publishing even for completely different records.

Using peeridlocker we only globally block access the first time a record is being published, and ONLY while creating
the locker for that particular peerID which is extremely short. Anytime after that, a block is held only for that peerid.
*/
package peeridlocker

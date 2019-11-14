..
..  Copyright (c) 2019 AT&T Intellectual Property.
..  Copyright (c) 2019 Nokia.
..
..  Licensed under the Creative Commons Attribution 4.0 International
..  Public License (the "License"); you may not use this file except
..  in compliance with the License. You may obtain a copy of the License at
..
..    https://creativecommons.org/licenses/by/4.0/
..
..  Unless required by applicable law or agreed to in writing, documentation
..  distributed under the License is distributed on an "AS IS" BASIS,
..  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
..
..  See the License for the specific language governing permissions and
..  limitations under the License.
..

Release-Notes
=============

This document provides the release notes for O-RAN SC Amber release of
ric-plt/streaming-protobufs.

.. contents::
   :depth: 3
   :local:



Version history
---------------

[0.2.2] - 2019-11-14

* Add mandatory documentation files.
* Fix Protobuf version number in x2ap_streaming.proto file.

[0.2.1] - 2019-10-16

* Fix protobuf definition of PLMN-Identity-EUTRA-5GC.

[0.2.0] - 2019-10-11

* Initial version.
* Implement Protobuf schema for following X2AP messages:

  * SGNB ADDITION REQUEST
  * SGNB ADDITION REQUEST ACKNOWLEDGE
  * SGNB ADDITION REQUEST REJECT
  * SGNB RECONFIGURATION COMPLETE
  * RRC TRANSFER
  * UE CONTEXT RELEASE
  * SGNB RELEASE REQUEST
  * SGNB RELEASE REQUEST ACKNOWLEDGE
  * SGNB RELEASE REQUIRED
  * SGNB RELEASE CONFIRM
  * SGNB MODIFICATION REQUEST
  * SGNB MODIFICATION REQUEST REJECT
  * SGNB MODIFICATION REQUEST ACKNOWLEDGE
  * SGNB MODIFICATION REQUIRED
  * SGNB MODIFICATION CONFIRM
  * SGNB MODIFICATION REFUSE
  * SN STATUS TRANSFER
  * SECONDARY RAT DATA USAGE REPORT

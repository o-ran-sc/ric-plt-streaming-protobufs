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

[6.0.0] - 2020-04-30

* Renamed RRC event trigger EventA3 configuration's a3_Threshold parameter
  name to a3_Offset and change parameter type to MeasTriggerQuantityOffset.

[5.0.0] - 2020-04-22

* Fixed CSI-MeasConfig configuration under ServingCellConfig to be optional.

[4.0.3] - 2020-04-15

* Added git extension to Golang streaming_protobufs module path name.

[4.0.2] - 2020-04-15

* Moved Golang source code generated from '.proto' files to new subdirectory
  go/streaming_protobufs to be in line with the Go package name of Protobufs.

[4.0.1] - 2020-04-15

* Add Golang module files.

[4.0.0] - 2020-04-14

* Changed scg-CellGroupConfig parameter type from Bytes to RRCReconfiguration
  to describe SCG configuration as a Protobuf structure. Moved all '.proto'
  files from streaming-protobufs root directory to proto subdirectory. Added
  Golang source code generated from '.proto' files into the repository.

[3.0.0] - 2020-02-14

* Changed all parameters string types to bytes type except parameters 'gNbID'
  and 'protobuf_revision' types in x2ap_streaming.proto file were still left
  a string type.

[2.0.0] - 2020-01-16

* Changed SgNB UE X2AP ID to be an optional IE in SGNB ADDITION REQUEST REJECT
  to be compatible with 3GPP 36.423 standard. API backward incompatible change.

[1.0.0] - 2020-01-14

* Fixed overlapping declaration error hit in Golang data structures, which
  were generated from SGNB CHANGE CONFIRM protobuf message. API backward
  incompatible change.

[0.2.3] - 2019-12-18

* Implement Protobuf schema for following X2AP messages:

  * SGNB CHANGE REQUIRED
  * SGNB CHANGE CONFIRM
  * SGNB CHANGE REFUSE
  * SGNB ACTIVITY NOTIFICATION
  * GNB STATUS INDICATION

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

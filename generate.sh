#!/bin/sh

cd resource

# Clean the garbage 
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/msg-header-2_0.xsd
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/xmldsig-core-schema.xsd
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/xlink.xsd
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/envelope.xsd
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/xml.xsd
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/wsse.xsd

sed 's/messageStatus\.type/messageStatusType/g' < msg-header-2_0.xsd > a
mv a msg-header-2_0.xsd

sed 's/non-empty-string/nomEmptyString/g' < msg-header-2_0.xsd > a
mv a msg-header-2_0.xsd

sed 's/severity\.type/severityType/g' < msg-header-2_0.xsd > a
mv a msg-header-2_0.xsd

sed 's/status\.type/statusType/g' < msg-header-2_0.xsd > a
mv a msg-header-2_0.xsd

sed 's/sequenceNumber\.type/sequenceNumberType/g' < msg-header-2_0.xsd > a
mv a msg-header-2_0.xsd

##### BFM
####curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/BargainFinderMaxRQ_v3-0-0.wsdl
####curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/shopping/BargainFinderMaxRQRS_v3-0-0.xsd
####
####wsdl2go < BargainFinderMaxRQ_v3-0-0.wsdl > ../../generated/sabreSoap/bfm.go
####sed 's/bargainfindermaxbinding/sabreSoap/g' < ../../generated/sabreSoap/bfm.go > ../../generated/sabreSoap/a
####mv ../../generated/sabreSoap/a ../../generated/sabreSoap/bfm.go

# SessionCreate
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/usg/SessionCreateRQ.wsdl
curl -O http://webservices.sabre.com/wsdl/sabreXML1.0.00/usg/SessionCreateRQRS.xsd
wsdl2go <  SessionCreateRQ.wsdl > sessionCreate.go
sed 's/sessioncreatesoapbinding/sabreSoap/g' < sessionCreate.go > a
mv a ../sabreSoap/sessionCreate.go
exit



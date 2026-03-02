package client

import "github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

// This file is used to split cases where the multiple parsers return the same data format

type GenericZone models.Zone // Used to retrieve all zones in a single request
type DomainZone models.Zone  // Used to retrieve only domain zones from the HTML body response
type ArpaZone models.Zone    // Used to retrieve only ARPA zones from the HTML body response

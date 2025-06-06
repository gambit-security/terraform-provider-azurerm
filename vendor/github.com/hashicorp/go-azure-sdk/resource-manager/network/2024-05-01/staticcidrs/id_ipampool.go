package staticcidrs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&IPamPoolId{})
}

var _ resourceids.ResourceId = &IPamPoolId{}

// IPamPoolId is a struct representing the Resource ID for a I Pam Pool
type IPamPoolId struct {
	SubscriptionId     string
	ResourceGroupName  string
	NetworkManagerName string
	IpamPoolName       string
}

// NewIPamPoolID returns a new IPamPoolId struct
func NewIPamPoolID(subscriptionId string, resourceGroupName string, networkManagerName string, ipamPoolName string) IPamPoolId {
	return IPamPoolId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		NetworkManagerName: networkManagerName,
		IpamPoolName:       ipamPoolName,
	}
}

// ParseIPamPoolID parses 'input' into a IPamPoolId
func ParseIPamPoolID(input string) (*IPamPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IPamPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IPamPoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIPamPoolIDInsensitively parses 'input' case-insensitively into a IPamPoolId
// note: this method should only be used for API response data and not user input
func ParseIPamPoolIDInsensitively(input string) (*IPamPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IPamPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IPamPoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IPamPoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetworkManagerName, ok = input.Parsed["networkManagerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "networkManagerName", input)
	}

	if id.IpamPoolName, ok = input.Parsed["ipamPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ipamPoolName", input)
	}

	return nil
}

// ValidateIPamPoolID checks that 'input' can be parsed as a I Pam Pool ID
func ValidateIPamPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIPamPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted I Pam Pool ID
func (id IPamPoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkManagers/%s/ipamPools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetworkManagerName, id.IpamPoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this I Pam Pool ID
func (id IPamPoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticNetworkManagers", "networkManagers", "networkManagers"),
		resourceids.UserSpecifiedSegment("networkManagerName", "networkManagerName"),
		resourceids.StaticSegment("staticIpamPools", "ipamPools", "ipamPools"),
		resourceids.UserSpecifiedSegment("ipamPoolName", "ipamPoolName"),
	}
}

// String returns a human-readable description of this I Pam Pool ID
func (id IPamPoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Network Manager Name: %q", id.NetworkManagerName),
		fmt.Sprintf("Ipam Pool Name: %q", id.IpamPoolName),
	}
	return fmt.Sprintf("I Pam Pool (%s)", strings.Join(components, "\n"))
}

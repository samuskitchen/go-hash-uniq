package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

type PreOrder struct {
	ClientID     string          `json:"client_id"`
	DocType      string          `json:"doc_type"`
	PurchaseDate string          `json:"purchase_date"`
	Items        []PreOrderItems `json:"items"`
}

type PreOrderItems struct {
	Material   string `json:"material"`
	Quantity   int    `json:"quantity"`
	SalesUnit  string `json:"sales_unit"`
	ItemNumber *int   `json:"item_number"`
}

// GenerateHash generates an SHA-256 hash based on the PreOrderDto fields: DocType, ClientId, and PurchaseDate,
// as well as a sorted concatenation of the Material, Quantity, and SalesUnit fields of each item in the Items slice.
//
// The items are sorted by Material, Quantity, and Sales Unit to ensure that their order does not affect
// the resulting hash. Only these three fields are considered for the hash generation.
//
// Returns a hex-encoded string representing the SHA-256 hash.
func (p *PreOrder) GenerateHash() string {
	var sb strings.Builder

	sb.WriteString(p.DocType)
	sb.WriteString(p.ClientID)
	sb.WriteString(p.PurchaseDate)

	// Sort items by Material, Quantity, and SalesUnit to ensure consistent order before generating the hash.
	sort.Slice(p.Items, func(i, j int) bool {
		if p.Items[i].Material != p.Items[j].Material {
			return p.Items[i].Material < p.Items[j].Material
		}

		if p.Items[i].Quantity != p.Items[j].Quantity {
			return p.Items[i].Quantity < p.Items[j].Quantity
		}

		return p.Items[i].SalesUnit < p.Items[j].SalesUnit
	})

	// Concatenate the Material, Quantity, and SalesUnit fields of each item.
	for _, item := range p.Items {
		sb.WriteString(fmt.Sprintf("%s%d%s", item.Material, item.Quantity, item.SalesUnit))
	}

	hash := sha256.Sum256([]byte(sb.String()))
	return hex.EncodeToString(hash[:])
}

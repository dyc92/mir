// Copyright 2019 Michael Li <dyc92@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package gin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gin Suite")
}

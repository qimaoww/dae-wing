/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package graphql

import (
	"github.com/qimaoww/dae-wing/graphql/service"
	"github.com/qimaoww/dae-wing/graphql/service/config"
	"github.com/qimaoww/dae-wing/graphql/service/config/global"
	"github.com/qimaoww/dae-wing/graphql/service/dns"
	"github.com/qimaoww/dae-wing/graphql/service/general"
	"github.com/qimaoww/dae-wing/graphql/service/group"
	"github.com/qimaoww/dae-wing/graphql/service/node"
	"github.com/qimaoww/dae-wing/graphql/service/routing"
	"github.com/qimaoww/dae-wing/graphql/service/subscription"
	"github.com/qimaoww/dae-wing/graphql/service/user"
)

type SchemaChain func() (string, error)

var schemaChains = []SchemaChain{
	general.Schema,
	config.Schema,
	global.Schema,
	group.Schema,
	routing.Schema,
	dns.Schema,
	service.Schema,
	node.Schema,
	subscription.Schema,
	user.Schema,
}

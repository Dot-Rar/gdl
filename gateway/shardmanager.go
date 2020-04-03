package gateway

import (
	"github.com/juju/ratelimit"
	"github.com/rxdn/gdl/cache"
	"github.com/rxdn/gdl/gateway/payloads/events"
	"github.com/rxdn/gdl/objects/user"
	"time"
)

type ShardManager struct {
	Token string

	GatewayBucket *ratelimit.Bucket

	TotalShards  int
	MinimumShard int // Inclusive
	MaximumShard int // Inclusive

	Shards map[int]*Shard

	EventBus *events.EventBus

	CacheFactory cache.CacheFactory

	Presence           user.UpdateStatus
	GuildSubscriptions bool
}

func NewShardManager(token string, shardOptions ShardOptions, cacheFactory cache.CacheFactory) ShardManager {
	manager := ShardManager{
		Token:         token,
		GatewayBucket: ratelimit.NewBucket(time.Second*6, 1),

		TotalShards:  shardOptions.Total,
		MinimumShard: shardOptions.Lowest,
		MaximumShard: shardOptions.Highest,

		EventBus:     events.NewEventBus(),
		CacheFactory: cacheFactory,
	}

	manager.Shards = make(map[int]*Shard)
	for i := shardOptions.Lowest; i <= shardOptions.Highest; i++ {
		shard := NewShard(&manager, token, i)
		manager.Shards[i] = &shard
	}

	RegisterCacheListeners(&manager)

	return manager
}

func (sm *ShardManager) Connect() {
	for _, shard := range sm.Shards {
		go shard.EnsureConnect()
	}
}

func (sm *ShardManager) RegisterListeners(listeners ...interface{}) {
	for _, listener := range listeners {
		sm.EventBus.RegisterListener(listener)
	}
}

func (sm *ShardManager) ShardForGuild(guildId uint64) *Shard {
	for _, shard := range sm.Shards {
		if shard.Cache.GetGuild(guildId) != nil {
			return shard
		}
	}

	return nil
}

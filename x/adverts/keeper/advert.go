package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"remap/x/adverts/types"
)

// GetAdvertCount get the total number of advert
func (k Keeper) GetAdvertCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdvertCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAdvertCount set the total number of advert
func (k Keeper) SetAdvertCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdvertCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAdvert appends a advert in the store with a new id and update the count
func (k Keeper) AppendAdvert(
	ctx sdk.Context,
	advert types.Advert,
) uint64 {
	// Create the advert
	count := k.GetAdvertCount(ctx)

	// Set the ID of the appended value
	advert.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdvertKey))
	appendedValue := k.cdc.MustMarshal(&advert)
	store.Set(GetAdvertIDBytes(advert.Id), appendedValue)

	// Update advert count
	k.SetAdvertCount(ctx, count+1)

	return count
}

// SetAdvert set a specific advert in the store
func (k Keeper) SetAdvert(ctx sdk.Context, advert types.Advert) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdvertKey))
	b := k.cdc.MustMarshal(&advert)
	store.Set(GetAdvertIDBytes(advert.Id), b)
}

// GetAdvert returns a advert from its id
func (k Keeper) GetAdvert(ctx sdk.Context, id uint64) (val types.Advert, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdvertKey))
	b := store.Get(GetAdvertIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAdvert removes a advert from the store
func (k Keeper) RemoveAdvert(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdvertKey))
	store.Delete(GetAdvertIDBytes(id))
}

// GetAllAdvert returns all advert
func (k Keeper) GetAllAdvert(ctx sdk.Context) (list []types.Advert) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdvertKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Advert
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAdvertIDBytes returns the byte representation of the ID
func GetAdvertIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAdvertIDFromBytes returns ID in uint64 format from a byte array
func GetAdvertIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

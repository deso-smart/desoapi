package desoapi

import (
	desoCore "github.com/deso-protocol/core/lib"
	"github.com/holiman/uint256"
)

type AuthorizeDerivedKeyRequest struct {
	// The original public key of the derived key owner.
	OwnerPublicKeyBase58Check string

	// The derived public key
	DerivedPublicKeyBase58Check string

	// The expiration block of the derived key pair.
	ExpirationBlock uint64

	// The signature of hash(derived key + expiration block) made by the owner.
	AccessSignature string

	// The intended operation on the derived key.
	DeleteKey bool

	// If we intend to sign this transaction with a derived key.
	DerivedKeySignature bool

	// No need to specify ProfileEntryResponse in each TransactionFee
	TransactionFees []TransactionFee

	MinFeeRateNanosPerKB uint64
}

type AuthorizeDerivedKeyResponse struct {
	SpendAmountNanos  uint64
	TotalInputNanos   uint64
	ChangeAmountNanos uint64
	FeeNanos          uint64
	Transaction       *desoCore.MsgDeSoTxn
	TransactionHex    string
	TxnHashHex        string
}

type TransactionFee struct {
	// PublicKeyBase58Check is the public key of the user who receives the fee.
	PublicKeyBase58Check string
	// ProfileEntryResponse is only non-nil when TransactionFees are retrieved through admin endpoints.
	// The ProfileEntryResponse is only used to display usernames and avatars in the admin dashboard and thus is
	// excluded in other places to reduce payload sizes and improve performance.
	ProfileEntryResponse *ProfileEntryResponse
	// AmountNanos is the amount PublicKeyBase58Check receives when this fee is incurred.
	AmountNanos uint64
}

type ProfileEntryResponse struct {
	// PublicKey is the key used by the user to sign for things and generally
	// verify her identity.
	PublicKeyBase58Check string
	Username             string
	Description          string
	IsHidden             bool
	IsReserved           bool
	IsVerified           bool
	Comments             []*PostEntryResponse
	Posts                []*PostEntryResponse
	// Creator coin fields
	CoinEntry *CoinEntryResponse

	// DAO Coin fields
	DAOCoinEntry *DAOCoinEntryResponse

	// Include current price for the frontend to display.
	CoinPriceDeSoNanos     uint64
	CoinPriceBitCloutNanos uint64 // Deprecated

	// Profiles of users that hold the coin + their balances.
	UsersThatHODL []*BalanceEntryResponse

	// If user is featured as a well known creator in the tutorial.
	IsFeaturedTutorialWellKnownCreator bool
	// If user is featured as an up and coming creator in the tutorial.
	// Note: a user should not be both featured as well known and up and coming
	IsFeaturedTutorialUpAndComingCreator bool
}

type PostEntryResponse struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	ParentStakeID              string
	Body                       string
	ImageURLs                  []string
	VideoURLs                  []string
	RepostedPostEntryResponse  *PostEntryResponse
	CreatorBasisPoints         uint64
	StakeMultipleBasisPoints   uint64
	TimestampNanos             uint64
	IsHidden                   bool
	ConfirmationBlockHeight    uint32
	InMempool                  bool
	// The profile associated with this post.
	ProfileEntryResponse *ProfileEntryResponse
	// The comments associated with this post.
	Comments     []*PostEntryResponse
	LikeCount    uint64
	DiamondCount uint64
	// Information about the reader's state w/regard to this post (e.g. if they liked it).
	PostEntryReaderState *desoCore.PostEntryReaderState
	InGlobalFeed         *bool
	InHotFeed            *bool
	// True if this post hash hex is pinned to the global feed.
	IsPinned *bool
	// PostExtraData stores an arbitrary map of attributes of a PostEntry
	PostExtraData    map[string]string
	CommentCount     uint64
	RepostCount      uint64
	QuoteRepostCount uint64
	// A list of parent posts for this post (ordered: root -> closest parent post).
	ParentPosts []*PostEntryResponse

	// NFT info.
	IsNFT                          bool
	NumNFTCopies                   uint64
	NumNFTCopiesForSale            uint64
	NumNFTCopiesBurned             uint64
	HasUnlockable                  bool
	NFTRoyaltyToCreatorBasisPoints uint64
	NFTRoyaltyToCoinBasisPoints    uint64
	// This map specifies royalties that should go to user's  other than the creator
	AdditionalDESORoyaltiesMap map[string]uint64
	// This map specifies royalties that should be add to creator coins other than the creator's coin.
	AdditionalCoinRoyaltiesMap map[string]uint64

	// Number of diamonds the sender gave this post. Only set when getting diamond posts.
	DiamondsFromSender uint64

	// Score given to this post by the hot feed go routine. Not always populated.
	HotnessScore   uint64
	PostMultiplier float64

	RecloutCount               uint64             // Deprecated
	QuoteRecloutCount          uint64             // Deprecated
	RecloutedPostEntryResponse *PostEntryResponse // Deprecated
}

type CoinEntryResponse struct {
	CreatorBasisPoints      uint64
	DeSoLockedNanos         uint64
	NumberOfHolders         uint64
	CoinsInCirculationNanos uint64
	CoinWatermarkNanos      uint64

	// Deprecated: Temporary to add support for BitCloutLockedNanos
	BitCloutLockedNanos uint64 // Deprecated
}

type DAOCoinEntryResponse struct {
	NumberOfHolders           uint64
	CoinsInCirculationNanos   uint256.Int
	MintingDisabled           bool
	TransferRestrictionStatus TransferRestrictionStatusString
}

type TransferRestrictionStatusString string

const (
	TransferRestrictionStatusStringUnrestricted            TransferRestrictionStatusString = "unrestricted"
	TransferRestrictionStatusStringProfileOwnerOnly        TransferRestrictionStatusString = "profile_owner_only"
	TransferRestrictionStatusStringDAOMembersOnly          TransferRestrictionStatusString = "dao_members_only"
	TransferRestrictionStatusStringPermanentlyUnrestricted TransferRestrictionStatusString = "permanently_unrestricted"
)

type BalanceEntryResponse struct {
	// The public keys are provided for the frontend
	HODLerPublicKeyBase58Check string
	// The public keys are provided for the frontend
	CreatorPublicKeyBase58Check string

	// Has the hodler purchased this creator's coin
	HasPurchased bool

	// How much this HODLer owns of a particular creator coin.
	BalanceNanos uint64

	// For simplicity, we create a new field for the uint256 balance for DAO coins
	BalanceNanosUint256 uint256.Int

	// The net effect of transactions in the mempool on a given BalanceEntry's BalanceNanos.
	// This is used by the frontend to convey info about mining.
	NetBalanceInMempool int64

	ProfileEntryResponse *ProfileEntryResponse
}

type GetExchangeRateResponse struct {
	// BTC
	SatoshisPerDeSoExchangeRate    uint64
	USDCentsPerBitcoinExchangeRate uint64

	// ETH
	NanosPerETHExchangeRate    uint64
	USDCentsPerETHExchangeRate uint64

	// DESO
	NanosSold                          uint64
	USDCentsPerDeSoExchangeRate        uint64
	USDCentsPerDeSoReserveExchangeRate uint64
	BuyDeSoFeeBasisPoints              uint64

	SatoshisPerBitCloutExchangeRate        uint64 // Deprecated
	USDCentsPerBitCloutExchangeRate        uint64 // Deprecated
	USDCentsPerBitCloutReserveExchangeRate uint64 // Deprecated
}

type BaseResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string

	// The information contained in the block’s header.
	Header *HeaderResponse

	Transactions []*TransactionResponse
}

type HeaderResponse struct {
	// The hash of the block that was queried.
	BlockHashHex string
	// Generally set to zero
	Version uint32
	// Hash of the previous block in the chain.
	PrevBlockHashHex string
	// The merkle root of all the transactions contained within the block.
	TransactionMerkleRootHex string
	// The unix timestamp (in seconds) specifying when this block was
	// mined.
	TstampSecs uint64
	// The height of the block this header corresponds to.
	Height uint64

	// The Nonce and ExtraNonce combine to give miners 128 bits of entropy
	Nonce      uint64
	ExtraNonce uint64
}

type TransactionResponse struct {
	// A string that uniquely identifies this transaction. This is a sha256 hash
	// of the transaction’s data encoded using base58 check encoding.
	TransactionIDBase58Check string
	// The raw hex of the transaction data. This can be fully-constructed from
	// the human-readable portions of this object.
	RawTransactionHex string
	// The inputs and outputs for this transaction.
	Inputs  []*InputResponse
	Outputs []*OutputResponse
	// The signature of the transaction in hex format.
	SignatureHex string
	// Will always be “0” for basic transfers
	TransactionType string

	// The hash of the block in which this transaction was mined. If the
	// transaction is unconfirmed, this field will be empty. To look up
	// how many confirmations a transaction has, simply plug this value
	// into the "block" endpoint.
	BlockHashHex string

	TransactionMetadata *desoCore.TransactionMetadata

	// The ExtraData added to this transaction
	ExtraData map[string]string
}

type InputResponse struct {
	TransactionIDBase58Check string
	Index                    int64
}

// OutputResponse ...
type OutputResponse struct {
	PublicKeyBase58Check string
	AmountNanos          uint64
}

type GetSingleProfileRequest struct {
	// When set, we return profiles starting at the given pubkey up to numEntriesToReturn.
	PublicKeyBase58Check string
	// When set, we return profiles starting at the given username up to numEntriesToReturn.
	Username string
	// When true, we don't log a 404 for missing profiles
	NoErrorOnMissing bool
}

type GetSingleProfileResponse struct {
	Profile       *ProfileEntryResponse
	IsBlacklisted bool
	IsGraylisted  bool
}

type SubmitTransactionRequest struct {
	TransactionHex string
}

type SubmitTransactionResponse struct {
	Transaction *desoCore.MsgDeSoTxn
	TxnHashHex  string

	// include the PostEntryResponse if a post was submitted
	PostEntryResponse *PostEntryResponse
}

type TransactionInfoRequest struct {
	// When set to true, the response simply contains all transactions in the
	// mempool with no filtering.
	IsMempool bool

	// A string that uniquely identifies this transaction. E.g. from a previous
	// call to “transfer-deso”. Ignored when PublicKeyBase58Check is set.
	TransactionIDBase58Check string

	// An DeSo public key encoded using base58 check encoding (starts
	// with “BC”) to get transaction IDs for. When set,
	// TransactionIDBase58Check is ignored.
	PublicKeyBase58Check string

	// Only return transaction IDs
	IDsOnly bool

	// Offset from which a page should be fetched
	LastTransactionIDBase58Check string

	// The last index of a transaction for a public key seen. If less than 0, it means we are not looking at
	// transactions in the database yet.
	LastPublicKeyTransactionIndex int64

	// Number of transactions to be returned
	Limit uint64
}

type TransactionInfoResponse struct {
	// Blank if successful. Otherwise, contains a description of the
	// error that occurred.
	Error string

	// The info for all transactions this public key is associated with from oldest
	// to newest.
	Transactions []*TransactionResponse

	// The hash of the last transaction
	LastTransactionIDBase58Check string

	// The last index of a transaction for a public key seen.
	LastPublicKeyTransactionIndex int64

	BalanceNanos uint64
}

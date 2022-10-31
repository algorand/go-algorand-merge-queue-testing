// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"encoding/json"
	"time"
)

const (
	Api_keyScopes = "api_key.Scopes"
)

// Defines values for AccountSigType.
const (
	AccountSigTypeLsig AccountSigType = "lsig"
	AccountSigTypeMsig AccountSigType = "msig"
	AccountSigTypeSig  AccountSigType = "sig"
)

// Defines values for AddressRole.
const (
	FreezeTarget AddressRole = "freeze-target"
	Receiver     AddressRole = "receiver"
	Sender       AddressRole = "sender"
)

// Defines values for Format.
const (
	Json    Format = "json"
	Msgpack Format = "msgpack"
)

// Defines values for SigType.
const (
	SigTypeLsig SigType = "lsig"
	SigTypeMsig SigType = "msig"
	SigTypeSig  SigType = "sig"
)

// Defines values for TxType.
const (
	Acfg   TxType = "acfg"
	Afrz   TxType = "afrz"
	Appl   TxType = "appl"
	Axfer  TxType = "axfer"
	Keyreg TxType = "keyreg"
	Pay    TxType = "pay"
	Stpf   TxType = "stpf"
)

// Defines values for TransactionProofResponseHashtype.
const (
	Sha256    TransactionProofResponseHashtype = "sha256"
	Sha512256 TransactionProofResponseHashtype = "sha512_256"
)

// Account Account information at a given round.
//
// Definition:
// data/basics/userBalance.go : AccountData
type Account struct {
	// Address the account public key
	Address string `json:"address"`

	// Amount \[algo\] total number of MicroAlgos in the account
	Amount uint64 `json:"amount"`

	// AmountWithoutPendingRewards specifies the amount of MicroAlgos in the account, without the pending rewards.
	AmountWithoutPendingRewards uint64 `json:"amount-without-pending-rewards"`

	// AppsLocalState \[appl\] applications local data stored in this account.
	//
	// Note the raw object uses `map[int] -> AppLocalState` for this type.
	AppsLocalState *[]ApplicationLocalState `json:"apps-local-state,omitempty"`

	// AppsTotalExtraPages \[teap\] the sum of all extra application program pages for this account.
	AppsTotalExtraPages *uint64 `json:"apps-total-extra-pages,omitempty"`

	// AppsTotalSchema Specifies maximums on the number of each type that may be stored.
	AppsTotalSchema *ApplicationStateSchema `json:"apps-total-schema,omitempty"`

	// Assets \[asset\] assets held by this account.
	//
	// Note the raw object uses `map[int] -> AssetHolding` for this type.
	Assets *[]AssetHolding `json:"assets,omitempty"`

	// AuthAddr \[spend\] the address against which signing should be checked. If empty, the address of the current account is used. This field can be updated in any transaction by setting the RekeyTo field.
	AuthAddr *string `json:"auth-addr,omitempty"`

	// CreatedApps \[appp\] parameters of applications created by this account including app global data.
	//
	// Note: the raw account uses `map[int] -> AppParams` for this type.
	CreatedApps *[]Application `json:"created-apps,omitempty"`

	// CreatedAssets \[apar\] parameters of assets created by this account.
	//
	// Note: the raw account uses `map[int] -> Asset` for this type.
	CreatedAssets *[]Asset `json:"created-assets,omitempty"`

	// MinBalance MicroAlgo balance required by the account.
	//
	// The requirement grows based on asset and application usage.
	MinBalance uint64 `json:"min-balance"`

	// Participation AccountParticipation describes the parameters used by this account in consensus protocol.
	Participation *AccountParticipation `json:"participation,omitempty"`

	// PendingRewards amount of MicroAlgos of pending rewards in this account.
	PendingRewards uint64 `json:"pending-rewards"`

	// RewardBase \[ebase\] used as part of the rewards computation. Only applicable to accounts which are participating.
	RewardBase *uint64 `json:"reward-base,omitempty"`

	// Rewards \[ern\] total rewards of MicroAlgos the account has received, including pending rewards.
	Rewards uint64 `json:"rewards"`

	// Round The round for which this information is relevant.
	Round uint64 `json:"round"`

	// SigType Indicates what type of signature is used by this account, must be one of:
	// * sig
	// * msig
	// * lsig
	SigType *AccountSigType `json:"sig-type,omitempty"`

	// Status \[onl\] delegation status of the account's MicroAlgos
	// * Offline - indicates that the associated account is delegated.
	// *  Online  - indicates that the associated account used as part of the delegation pool.
	// *   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.
	Status string `json:"status"`

	// TotalAppsOptedIn The count of all applications that have been opted in, equivalent to the count of application local data (AppLocalState objects) stored in this account.
	TotalAppsOptedIn uint64 `json:"total-apps-opted-in"`

	// TotalAssetsOptedIn The count of all assets that have been opted in, equivalent to the count of AssetHolding objects held by this account.
	TotalAssetsOptedIn uint64 `json:"total-assets-opted-in"`

	// TotalBoxBytes \[tbxb\] The total number of bytes used by this account's app's box keys and values.
	TotalBoxBytes *uint64 `json:"total-box-bytes,omitempty"`

	// TotalBoxes \[tbx\] The number of existing boxes created by this account's app.
	TotalBoxes *uint64 `json:"total-boxes,omitempty"`

	// TotalCreatedApps The count of all apps (AppParams objects) created by this account.
	TotalCreatedApps uint64 `json:"total-created-apps"`

	// TotalCreatedAssets The count of all assets (AssetParams objects) created by this account.
	TotalCreatedAssets uint64 `json:"total-created-assets"`
}

// AccountSigType Indicates what type of signature is used by this account, must be one of:
// * sig
// * msig
// * lsig
type AccountSigType string

// AccountParticipation AccountParticipation describes the parameters used by this account in consensus protocol.
type AccountParticipation struct {
	// SelectionParticipationKey \[sel\] Selection public key (if any) currently registered for this round.
	SelectionParticipationKey []byte `json:"selection-participation-key"`

	// StateProofKey \[stprf\] Root of the state proof key (if any)
	StateProofKey *[]byte `json:"state-proof-key,omitempty"`

	// VoteFirstValid \[voteFst\] First round for which this participation is valid.
	VoteFirstValid uint64 `json:"vote-first-valid"`

	// VoteKeyDilution \[voteKD\] Number of subkeys in each batch of participation keys.
	VoteKeyDilution uint64 `json:"vote-key-dilution"`

	// VoteLastValid \[voteLst\] Last round for which this participation is valid.
	VoteLastValid uint64 `json:"vote-last-valid"`

	// VoteParticipationKey \[vote\] root participation public key (if any) currently registered for this round.
	VoteParticipationKey []byte `json:"vote-participation-key"`
}

// AccountStateDelta Application state delta.
type AccountStateDelta struct {
	Address string `json:"address"`

	// Delta Application state delta.
	Delta StateDelta `json:"delta"`
}

// Application Application index and its parameters
type Application struct {
	// Id \[appidx\] application index.
	Id uint64 `json:"id"`

	// Params Stores the global information associated with an application.
	Params ApplicationParams `json:"params"`
}

// ApplicationLocalState Stores local state associated with an application.
type ApplicationLocalState struct {
	// Id The application which this local state is for.
	Id uint64 `json:"id"`

	// KeyValue Represents a key-value store for use in an application.
	KeyValue *TealKeyValueStore `json:"key-value,omitempty"`

	// Schema Specifies maximums on the number of each type that may be stored.
	Schema ApplicationStateSchema `json:"schema"`
}

// ApplicationParams Stores the global information associated with an application.
type ApplicationParams struct {
	// ApprovalProgram \[approv\] approval program.
	ApprovalProgram []byte `json:"approval-program"`

	// ClearStateProgram \[clearp\] approval program.
	ClearStateProgram []byte `json:"clear-state-program"`

	// Creator The address that created this application. This is the address where the parameters and global state for this application can be found.
	Creator string `json:"creator"`

	// ExtraProgramPages \[epp\] the amount of extra program pages available to this app.
	ExtraProgramPages *uint64 `json:"extra-program-pages,omitempty"`

	// GlobalState Represents a key-value store for use in an application.
	GlobalState *TealKeyValueStore `json:"global-state,omitempty"`

	// GlobalStateSchema Specifies maximums on the number of each type that may be stored.
	GlobalStateSchema *ApplicationStateSchema `json:"global-state-schema,omitempty"`

	// LocalStateSchema Specifies maximums on the number of each type that may be stored.
	LocalStateSchema *ApplicationStateSchema `json:"local-state-schema,omitempty"`
}

// ApplicationStateSchema Specifies maximums on the number of each type that may be stored.
type ApplicationStateSchema struct {
	// NumByteSlice \[nbs\] num of byte slices.
	NumByteSlice uint64 `json:"num-byte-slice"`

	// NumUint \[nui\] num of uints.
	NumUint uint64 `json:"num-uint"`
}

// Asset Specifies both the unique identifier and the parameters for an asset
type Asset struct {
	// Index unique asset identifier
	Index uint64 `json:"index"`

	// Params AssetParams specifies the parameters for an asset.
	//
	// \[apar\] when part of an AssetConfig transaction.
	//
	// Definition:
	// data/transactions/asset.go : AssetParams
	Params AssetParams `json:"params"`
}

// AssetHolding Describes an asset held by an account.
//
// Definition:
// data/basics/userBalance.go : AssetHolding
type AssetHolding struct {
	// Amount \[a\] number of units held.
	Amount uint64 `json:"amount"`

	// AssetId Asset ID of the holding.
	AssetID uint64 `json:"asset-id"`

	// IsFrozen \[f\] whether or not the holding is frozen.
	IsFrozen bool `json:"is-frozen"`
}

// AssetParams AssetParams specifies the parameters for an asset.
//
// \[apar\] when part of an AssetConfig transaction.
//
// Definition:
// data/transactions/asset.go : AssetParams
type AssetParams struct {
	// Clawback \[c\] Address of account used to clawback holdings of this asset.  If empty, clawback is not permitted.
	Clawback *string `json:"clawback,omitempty"`

	// Creator The address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.
	Creator string `json:"creator"`

	// Decimals \[dc\] The number of digits to use after the decimal point when displaying this asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, and so on. This value must be between 0 and 19 (inclusive).
	Decimals uint64 `json:"decimals"`

	// DefaultFrozen \[df\] Whether holdings of this asset are frozen by default.
	DefaultFrozen *bool `json:"default-frozen,omitempty"`

	// Freeze \[f\] Address of account used to freeze holdings of this asset.  If empty, freezing is not permitted.
	Freeze *string `json:"freeze,omitempty"`

	// Manager \[m\] Address of account used to manage the keys of this asset and to destroy it.
	Manager *string `json:"manager,omitempty"`

	// MetadataHash \[am\] A commitment to some unspecified asset metadata. The format of this metadata is up to the application.
	MetadataHash *[]byte `json:"metadata-hash,omitempty"`

	// Name \[an\] Name of this asset, as supplied by the creator. Included only when the asset name is composed of printable utf-8 characters.
	Name *string `json:"name,omitempty"`

	// NameB64 Base64 encoded name of this asset, as supplied by the creator.
	NameB64 *[]byte `json:"name-b64,omitempty"`

	// Reserve \[r\] Address of account holding reserve (non-minted) units of this asset.
	Reserve *string `json:"reserve,omitempty"`

	// Total \[t\] The total number of units of this asset.
	Total uint64 `json:"total"`

	// UnitName \[un\] Name of a unit of this asset, as supplied by the creator. Included only when the name of a unit of this asset is composed of printable utf-8 characters.
	UnitName *string `json:"unit-name,omitempty"`

	// UnitNameB64 Base64 encoded name of a unit of this asset, as supplied by the creator.
	UnitNameB64 *[]byte `json:"unit-name-b64,omitempty"`

	// Url \[au\] URL where more information about the asset can be retrieved. Included only when the URL is composed of printable utf-8 characters.
	Url *string `json:"url,omitempty"`

	// UrlB64 Base64 encoded URL where more information about the asset can be retrieved.
	UrlB64 *[]byte `json:"url-b64,omitempty"`
}

// Box Box name and its content.
type Box struct {
	// Name \[name\] box name, base64 encoded
	Name []byte `json:"name"`

	// Value \[value\] box value, base64 encoded.
	Value []byte `json:"value"`
}

// BoxDescriptor Box descriptor describes a Box.
type BoxDescriptor struct {
	// Name Base64 encoded box name
	Name []byte `json:"name"`
}

// BuildVersion defines model for BuildVersion.
type BuildVersion struct {
	Branch      string `json:"branch"`
	BuildNumber uint64 `json:"build_number"`
	Channel     string `json:"channel"`
	CommitHash  string `json:"commit_hash"`
	Major       uint64 `json:"major"`
	Minor       uint64 `json:"minor"`
}

// DryrunRequest Request data type for dryrun endpoint. Given the Transactions and simulated ledger state upload, run TEAL scripts and return debugging information.
type DryrunRequest struct {
	Accounts []Account     `json:"accounts"`
	Apps     []Application `json:"apps"`

	// LatestTimestamp LatestTimestamp is available to some TEAL scripts. Defaults to the latest confirmed timestamp this algod is attached to.
	LatestTimestamp uint64 `json:"latest-timestamp"`

	// ProtocolVersion ProtocolVersion specifies a specific version string to operate under, otherwise whatever the current protocol of the network this algod is running in.
	ProtocolVersion string `json:"protocol-version"`

	// Round Round is available to some TEAL scripts. Defaults to the current round on the network this algod is attached to.
	Round   uint64            `json:"round"`
	Sources []DryrunSource    `json:"sources"`
	Txns    []json.RawMessage `json:"txns"`
}

// DryrunSource DryrunSource is TEAL source text that gets uploaded, compiled, and inserted into transactions or application state.
type DryrunSource struct {
	AppIndex uint64 `json:"app-index"`

	// FieldName FieldName is what kind of sources this is. If lsig then it goes into the transactions[this.TxnIndex].LogicSig. If approv or clearp it goes into the Approval Program or Clear State Program of application[this.AppIndex].
	FieldName string `json:"field-name"`
	Source    string `json:"source"`
	TxnIndex  uint64 `json:"txn-index"`
}

// DryrunState Stores the TEAL eval step data
type DryrunState struct {
	// Error Evaluation error if any
	Error *string `json:"error,omitempty"`

	// Line Line number
	Line uint64 `json:"line"`

	// Pc Program counter
	Pc      uint64       `json:"pc"`
	Scratch *[]TealValue `json:"scratch,omitempty"`
	Stack   []TealValue  `json:"stack"`
}

// DryrunTxnResult DryrunTxnResult contains any LogicSig or ApplicationCall program debug information and state updates from a dryrun.
type DryrunTxnResult struct {
	AppCallMessages *[]string      `json:"app-call-messages,omitempty"`
	AppCallTrace    *[]DryrunState `json:"app-call-trace,omitempty"`

	// BudgetAdded Budget added during execution of app call transaction.
	BudgetAdded *uint64 `json:"budget-added,omitempty"`

	// BudgetConsumed Budget consumed during execution of app call transaction.
	BudgetConsumed *uint64 `json:"budget-consumed,omitempty"`

	// Cost Net cost of app execution. Field is DEPRECATED and is subject for removal. Instead, use `budget-added` and `budget-consumed.
	Cost *uint64 `json:"cost,omitempty"`

	// Disassembly Disassembled program line by line.
	Disassembly []string `json:"disassembly"`

	// GlobalDelta Application state delta.
	GlobalDelta *StateDelta          `json:"global-delta,omitempty"`
	LocalDeltas *[]AccountStateDelta `json:"local-deltas,omitempty"`

	// LogicSigDisassembly Disassembled lsig program line by line.
	LogicSigDisassembly *[]string      `json:"logic-sig-disassembly,omitempty"`
	LogicSigMessages    *[]string      `json:"logic-sig-messages,omitempty"`
	LogicSigTrace       *[]DryrunState `json:"logic-sig-trace,omitempty"`
	Logs                *[][]byte      `json:"logs,omitempty"`
}

// ErrorResponse An error response with optional data field.
type ErrorResponse struct {
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
}

// EvalDelta Represents a TEAL value delta.
type EvalDelta struct {
	// Action \[at\] delta action.
	Action uint64 `json:"action"`

	// Bytes \[bs\] bytes value.
	Bytes *string `json:"bytes,omitempty"`

	// Uint \[ui\] uint value.
	Uint *uint64 `json:"uint,omitempty"`
}

// EvalDeltaKeyValue Key-value pairs for StateDelta.
type EvalDeltaKeyValue struct {
	Key string `json:"key"`

	// Value Represents a TEAL value delta.
	Value EvalDelta `json:"value"`
}

// LightBlockHeaderProof Proof of membership and position of a light block header.
type LightBlockHeaderProof struct {
	// Index The index of the light block header in the vector commitment tree
	Index uint64 `json:"index"`

	// Proof The encoded proof.
	Proof []byte `json:"proof"`

	// Treedepth Represents the depth of the tree that is being proven, i.e. the number of edges from a leaf to the root.
	Treedepth uint64 `json:"treedepth"`
}

// ParticipationKey Represents a participation key used by the node.
type ParticipationKey struct {
	// Address Address the key was generated for.
	Address string `json:"address"`

	// EffectiveFirstValid When registered, this is the first round it may be used.
	EffectiveFirstValid *uint64 `json:"effective-first-valid,omitempty"`

	// EffectiveLastValid When registered, this is the last round it may be used.
	EffectiveLastValid *uint64 `json:"effective-last-valid,omitempty"`

	// Id The key's ParticipationID.
	Id string `json:"id"`

	// Key AccountParticipation describes the parameters used by this account in consensus protocol.
	Key AccountParticipation `json:"key"`

	// LastBlockProposal Round when this key was last used to propose a block.
	LastBlockProposal *uint64 `json:"last-block-proposal,omitempty"`

	// LastStateProof Round when this key was last used to generate a state proof.
	LastStateProof *uint64 `json:"last-state-proof,omitempty"`

	// LastVote Round when this key was last used to vote.
	LastVote *uint64 `json:"last-vote,omitempty"`
}

// PendingTransactionResponse Details about a pending transaction. If the transaction was recently confirmed, includes confirmation details like the round and reward details.
type PendingTransactionResponse struct {
	// ApplicationIndex The application index if the transaction was found and it created an application.
	ApplicationIndex *uint64 `json:"application-index,omitempty"`

	// AssetClosingAmount The number of the asset's unit that were transferred to the close-to address.
	AssetClosingAmount *uint64 `json:"asset-closing-amount,omitempty"`

	// AssetIndex The asset index if the transaction was found and it created an asset.
	AssetIndex *uint64 `json:"asset-index,omitempty"`

	// CloseRewards Rewards in microalgos applied to the close remainder to account.
	CloseRewards *uint64 `json:"close-rewards,omitempty"`

	// ClosingAmount Closing amount for the transaction.
	ClosingAmount *uint64 `json:"closing-amount,omitempty"`

	// ConfirmedRound The round where this transaction was confirmed, if present.
	ConfirmedRound *uint64 `json:"confirmed-round,omitempty"`

	// GlobalStateDelta Application state delta.
	GlobalStateDelta *StateDelta `json:"global-state-delta,omitempty"`

	// InnerTxns Inner transactions produced by application execution.
	InnerTxns *[]PendingTransactionResponse `json:"inner-txns,omitempty"`

	// LocalStateDelta \[ld\] Local state key/value changes for the application being executed by this transaction.
	LocalStateDelta *[]AccountStateDelta `json:"local-state-delta,omitempty"`

	// Logs \[lg\] Logs for the application being executed by this transaction.
	Logs *[][]byte `json:"logs,omitempty"`

	// PoolError Indicates that the transaction was kicked out of this node's transaction pool (and specifies why that happened).  An empty string indicates the transaction wasn't kicked out of this node's txpool due to an error.
	PoolError string `json:"pool-error"`

	// ReceiverRewards Rewards in microalgos applied to the receiver account.
	ReceiverRewards *uint64 `json:"receiver-rewards,omitempty"`

	// SenderRewards Rewards in microalgos applied to the sender account.
	SenderRewards *uint64 `json:"sender-rewards,omitempty"`

	// Txn The raw signed transaction.
	Txn map[string]interface{} `json:"txn"`
}

// StateDelta Application state delta.
type StateDelta = []EvalDeltaKeyValue

// StateProof Represents a state proof and its corresponding message
type StateProof struct {
	// Message Represents the message that the state proofs are attesting to.
	Message StateProofMessage `json:"Message"`

	// StateProof The encoded StateProof for the message.
	StateProof []byte `json:"StateProof"`
}

// StateProofMessage Represents the message that the state proofs are attesting to.
type StateProofMessage struct {
	// BlockHeadersCommitment The vector commitment root on all light block headers within a state proof interval.
	BlockHeadersCommitment []byte `json:"BlockHeadersCommitment"`

	// FirstAttestedRound The first round the message attests to.
	FirstAttestedRound uint64 `json:"FirstAttestedRound"`

	// LastAttestedRound The last round the message attests to.
	LastAttestedRound uint64 `json:"LastAttestedRound"`

	// LnProvenWeight An integer value representing the natural log of the proven weight with 16 bits of precision. This value would be used to verify the next state proof.
	LnProvenWeight uint64 `json:"LnProvenWeight"`

	// VotersCommitment The vector commitment root of the top N accounts to sign the next StateProof.
	VotersCommitment []byte `json:"VotersCommitment"`
}

// TealKeyValue Represents a key-value pair in an application store.
type TealKeyValue struct {
	Key string `json:"key"`

	// Value Represents a TEAL value.
	Value TealValue `json:"value"`
}

// TealKeyValueStore Represents a key-value store for use in an application.
type TealKeyValueStore = []TealKeyValue

// TealValue Represents a TEAL value.
type TealValue struct {
	// Bytes \[tb\] bytes value.
	Bytes string `json:"bytes"`

	// Type \[tt\] value type. Value `1` refers to **bytes**, value `2` refers to **uint**
	Type uint64 `json:"type"`

	// Uint \[ui\] uint value.
	Uint uint64 `json:"uint"`
}

// Version algod version information.
type Version struct {
	Build          BuildVersion `json:"build"`
	GenesisHashB64 []byte       `json:"genesis_hash_b64"`
	GenesisId      string       `json:"genesis_id"`
	Versions       []string     `json:"versions"`
}

// AccountID defines model for account-id.
type AccountID = string

// Address defines model for address.
type Address = string

// AddressRole defines model for address-role.
type AddressRole string

// AfterTime defines model for after-time.
type AfterTime = time.Time

// AssetID defines model for asset-id.
type AssetID uint64

// BeforeTime defines model for before-time.
type BeforeTime = time.Time

// Catchpoint defines model for catchpoint.
type Catchpoint = string

// CurrencyGreaterThan defines model for currency-greater-than.
type CurrencyGreaterThan uint64

// CurrencyLessThan defines model for currency-less-than.
type CurrencyLessThan uint64

// ExcludeCloseTo defines model for exclude-close-to.
type ExcludeCloseTo = bool

// Format defines model for format.
type Format string

// Limit defines model for limit.
type Limit uint64

// Max defines model for max.
type Max uint64

// MaxRound defines model for max-round.
type MaxRound uint64

// MinRound defines model for min-round.
type MinRound uint64

// Next defines model for next.
type Next = string

// NotePrefix defines model for note-prefix.
type NotePrefix = string

// Round defines model for round.
type Round uint64

// RoundNumber defines model for round-number.
type RoundNumber uint64

// SigType defines model for sig-type.
type SigType string

// TxID defines model for tx-id.
type TxID = string

// TxType defines model for tx-type.
type TxType string

// AccountApplicationResponse defines model for AccountApplicationResponse.
type AccountApplicationResponse struct {
	// AppLocalState Stores local state associated with an application.
	AppLocalState *ApplicationLocalState `json:"app-local-state,omitempty"`

	// CreatedApp Stores the global information associated with an application.
	CreatedApp *ApplicationParams `json:"created-app,omitempty"`

	// Round The round for which this information is relevant.
	Round uint64 `json:"round"`
}

// AccountAssetResponse defines model for AccountAssetResponse.
type AccountAssetResponse struct {
	// AssetHolding Describes an asset held by an account.
	//
	// Definition:
	// data/basics/userBalance.go : AssetHolding
	AssetHolding *AssetHolding `json:"asset-holding,omitempty"`

	// CreatedAsset AssetParams specifies the parameters for an asset.
	//
	// \[apar\] when part of an AssetConfig transaction.
	//
	// Definition:
	// data/transactions/asset.go : AssetParams
	CreatedAsset *AssetParams `json:"created-asset,omitempty"`

	// Round The round for which this information is relevant.
	Round uint64 `json:"round"`
}

// AccountResponse Account information at a given round.
//
// Definition:
// data/basics/userBalance.go : AccountData
type AccountResponse = Account

// ApplicationResponse Application index and its parameters
type ApplicationResponse = Application

// AssetResponse Specifies both the unique identifier and the parameters for an asset
type AssetResponse = Asset

// BlockHashResponse defines model for BlockHashResponse.
type BlockHashResponse struct {
	// BlockHash Block header hash.
	BlockHash string `json:"blockHash"`
}

// BlockResponse defines model for BlockResponse.
type BlockResponse struct {
	// Block Block header data.
	Block map[string]interface{} `json:"block"`

	// Cert Optional certificate object. This is only included when the format is set to message pack.
	Cert *map[string]interface{} `json:"cert,omitempty"`
}

// BoxResponse Box name and its content.
type BoxResponse = Box

// BoxesResponse defines model for BoxesResponse.
type BoxesResponse struct {
	Boxes []BoxDescriptor `json:"boxes"`
}

// CatchpointAbortResponse An catchpoint abort response.
type CatchpointAbortResponse struct {
	// CatchupMessage Catchup abort response string
	CatchupMessage string `json:"catchup-message"`
}

// CatchpointStartResponse An catchpoint start response.
type CatchpointStartResponse struct {
	// CatchupMessage Catchup start response string
	CatchupMessage string `json:"catchup-message"`
}

// CompileResponse defines model for CompileResponse.
type CompileResponse struct {
	// Hash base32 SHA512_256 of program bytes (Address style)
	Hash string `json:"hash"`

	// Result base64 encoded program bytes
	Result string `json:"result"`

	// Sourcemap JSON of the source map
	Sourcemap *map[string]interface{} `json:"sourcemap,omitempty"`
}

// DisassembleResponse defines model for DisassembleResponse.
type DisassembleResponse struct {
	// Result disassembled Teal code
	Result string `json:"result"`
}

// DryrunResponse defines model for DryrunResponse.
type DryrunResponse struct {
	Error string `json:"error"`

	// ProtocolVersion Protocol version is the protocol version Dryrun was operated under.
	ProtocolVersion string            `json:"protocol-version"`
	Txns            []DryrunTxnResult `json:"txns"`
}

// LightBlockHeaderProofResponse Proof of membership and position of a light block header.
type LightBlockHeaderProofResponse = LightBlockHeaderProof

// NodeStatusResponse NodeStatus contains the information about a node status
type NodeStatusResponse struct {
	// Catchpoint The current catchpoint that is being caught up to
	Catchpoint *string `json:"catchpoint,omitempty"`

	// CatchpointAcquiredBlocks The number of blocks that have already been obtained by the node as part of the catchup
	CatchpointAcquiredBlocks *uint64 `json:"catchpoint-acquired-blocks,omitempty"`

	// CatchpointProcessedAccounts The number of accounts from the current catchpoint that have been processed so far as part of the catchup
	CatchpointProcessedAccounts *uint64 `json:"catchpoint-processed-accounts,omitempty"`

	// CatchpointTotalAccounts The total number of accounts included in the current catchpoint
	CatchpointTotalAccounts *uint64 `json:"catchpoint-total-accounts,omitempty"`

	// CatchpointTotalBlocks The total number of blocks that are required to complete the current catchpoint catchup
	CatchpointTotalBlocks *uint64 `json:"catchpoint-total-blocks,omitempty"`

	// CatchpointVerifiedAccounts The number of accounts from the current catchpoint that have been verified so far as part of the catchup
	CatchpointVerifiedAccounts *uint64 `json:"catchpoint-verified-accounts,omitempty"`

	// CatchupTime CatchupTime in nanoseconds
	CatchupTime uint64 `json:"catchup-time"`

	// LastCatchpoint The last catchpoint seen by the node
	LastCatchpoint *string `json:"last-catchpoint,omitempty"`

	// LastRound LastRound indicates the last round seen
	LastRound uint64 `json:"last-round"`

	// LastVersion LastVersion indicates the last consensus version supported
	LastVersion string `json:"last-version"`

	// NextVersion NextVersion of consensus protocol to use
	NextVersion string `json:"next-version"`

	// NextVersionRound NextVersionRound is the round at which the next consensus version will apply
	NextVersionRound uint64 `json:"next-version-round"`

	// NextVersionSupported NextVersionSupported indicates whether the next consensus version is supported by this node
	NextVersionSupported bool `json:"next-version-supported"`

	// StoppedAtUnsupportedRound StoppedAtUnsupportedRound indicates that the node does not support the new rounds and has stopped making progress
	StoppedAtUnsupportedRound bool `json:"stopped-at-unsupported-round"`

	// TimeSinceLastRound TimeSinceLastRound in nanoseconds
	TimeSinceLastRound uint64 `json:"time-since-last-round"`
}

// ParticipationKeyResponse Represents a participation key used by the node.
type ParticipationKeyResponse = ParticipationKey

// ParticipationKeysResponse defines model for ParticipationKeysResponse.
type ParticipationKeysResponse = []ParticipationKey

// PendingTransactionsResponse PendingTransactions is an array of signed transactions exactly as they were submitted.
type PendingTransactionsResponse struct {
	// TopTransactions An array of signed transaction objects.
	TopTransactions []map[string]interface{} `json:"top-transactions"`

	// TotalTransactions Total number of transactions in the pool.
	TotalTransactions uint64 `json:"total-transactions"`
}

// PostParticipationResponse defines model for PostParticipationResponse.
type PostParticipationResponse struct {
	// PartId encoding of the participation ID.
	PartId string `json:"partId"`
}

// PostTransactionsResponse defines model for PostTransactionsResponse.
type PostTransactionsResponse struct {
	// TxId encoding of the transaction hash.
	TxId string `json:"txId"`
}

// StateProofResponse Represents a state proof and its corresponding message
type StateProofResponse = StateProof

// SupplyResponse Supply represents the current supply of MicroAlgos in the system
type SupplyResponse struct {
	// CurrentRound Round
	CurrentRound uint64 `json:"current_round"`

	// OnlineMoney OnlineMoney
	OnlineMoney uint64 `json:"online-money"`

	// TotalMoney TotalMoney
	TotalMoney uint64 `json:"total-money"`
}

// TransactionParametersResponse TransactionParams contains the parameters that help a client construct
// a new transaction.
type TransactionParametersResponse struct {
	// ConsensusVersion ConsensusVersion indicates the consensus protocol version
	// as of LastRound.
	ConsensusVersion string `json:"consensus-version"`

	// Fee Fee is the suggested transaction fee
	// Fee is in units of micro-Algos per byte.
	// Fee may fall to zero but transactions must still have a fee of
	// at least MinTxnFee for the current network protocol.
	Fee uint64 `json:"fee"`

	// GenesisHash GenesisHash is the hash of the genesis block.
	GenesisHash []byte `json:"genesis-hash"`

	// GenesisId GenesisID is an ID listed in the genesis block.
	GenesisId string `json:"genesis-id"`

	// LastRound LastRound indicates the last round seen
	LastRound uint64 `json:"last-round"`

	// MinFee The minimum transaction fee (not per byte) required for the
	// txn to validate for the current network protocol.
	MinFee uint64 `json:"min-fee"`
}

// TransactionProofResponse defines model for TransactionProofResponse.
type TransactionProofResponse struct {
	// Hashtype The type of hash function used to create the proof, must be one of:
	// * sha512_256
	// * sha256
	Hashtype TransactionProofResponseHashtype `json:"hashtype"`

	// Idx Index of the transaction in the block's payset.
	Idx uint64 `json:"idx"`

	// Proof Proof of transaction membership.
	Proof []byte `json:"proof"`

	// Stibhash Hash of SignedTxnInBlock for verifying proof.
	Stibhash []byte `json:"stibhash"`

	// Treedepth Represents the depth of the tree that is being proven, i.e. the number of edges from a leaf to the root.
	Treedepth uint64 `json:"treedepth"`
}

// TransactionProofResponseHashtype The type of hash function used to create the proof, must be one of:
// * sha512_256
// * sha256
type TransactionProofResponseHashtype string

// VersionsResponse algod version information.
type VersionsResponse = Version

// ShutdownNodeParams defines parameters for ShutdownNode.
type ShutdownNodeParams struct {
	Timeout *uint64 `form:"timeout,omitempty" json:"timeout,omitempty"`
}

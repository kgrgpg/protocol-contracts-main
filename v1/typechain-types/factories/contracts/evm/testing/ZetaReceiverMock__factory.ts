/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  ZetaReceiverMock,
  ZetaReceiverMockInterface,
} from "../../../../contracts/evm/testing/ZetaReceiverMock";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "destinationAddress",
        type: "address",
      },
    ],
    name: "MockOnZetaMessage",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "zetaTxSenderAddress",
        type: "address",
      },
    ],
    name: "MockOnZetaRevert",
    type: "event",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "bytes",
            name: "zetaTxSenderAddress",
            type: "bytes",
          },
          {
            internalType: "uint256",
            name: "sourceChainId",
            type: "uint256",
          },
          {
            internalType: "address",
            name: "destinationAddress",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "zetaValue",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "message",
            type: "bytes",
          },
        ],
        internalType: "struct ZetaInterfaces.ZetaMessage",
        name: "zetaMessage",
        type: "tuple",
      },
    ],
    name: "onZetaMessage",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "zetaTxSenderAddress",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "sourceChainId",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "destinationAddress",
            type: "bytes",
          },
          {
            internalType: "uint256",
            name: "destinationChainId",
            type: "uint256",
          },
          {
            internalType: "uint256",
            name: "remainingZetaValue",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "message",
            type: "bytes",
          },
        ],
        internalType: "struct ZetaInterfaces.ZetaRevert",
        name: "zetaRevert",
        type: "tuple",
      },
    ],
    name: "onZetaRevert",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b506102d5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633749c51a1461003b5780633ff0693c14610057575b600080fd5b6100556004803603810190610050919061018b565b610073565b005b610071600480360381019061006c91906101d4565b6100bf565b005b7f72a301dee3abcbe15615f3e253229bf4b4f508460603d674991c9a777b833c6e8160400160208101906100a7919061015e565b6040516100b4919061022c565b60405180910390a150565b7f53bd04e26f94f13ff43da96839541821041c309c6f624712192cbe3a2d133cc48160000160208101906100f3919061015e565b604051610100919061022c565b60405180910390a150565b60008135905061011a81610288565b92915050565b600060a0828403121561013657610135610279565b5b81905092915050565b600060c0828403121561015557610154610279565b5b81905092915050565b60006020828403121561017457610173610283565b5b60006101828482850161010b565b91505092915050565b6000602082840312156101a1576101a0610283565b5b600082013567ffffffffffffffff8111156101bf576101be61027e565b5b6101cb84828501610120565b91505092915050565b6000602082840312156101ea576101e9610283565b5b600082013567ffffffffffffffff8111156102085761020761027e565b5b6102148482850161013f565b91505092915050565b61022681610247565b82525050565b6000602082019050610241600083018461021d565b92915050565b600061025282610259565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b600080fd5b600080fd5b61029181610247565b811461029c57600080fd5b5056fea26469706673582212206d3535c1f777f93a6d105378513b52ef987ededd64b301e5cc1914be4e5ba14564736f6c63430008070033";

type ZetaReceiverMockConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ZetaReceiverMockConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ZetaReceiverMock__factory extends ContractFactory {
  constructor(...args: ZetaReceiverMockConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ZetaReceiverMock> {
    return super.deploy(overrides || {}) as Promise<ZetaReceiverMock>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): ZetaReceiverMock {
    return super.attach(address) as ZetaReceiverMock;
  }
  override connect(signer: Signer): ZetaReceiverMock__factory {
    return super.connect(signer) as ZetaReceiverMock__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZetaReceiverMockInterface {
    return new utils.Interface(_abi) as ZetaReceiverMockInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ZetaReceiverMock {
    return new Contract(address, _abi, signerOrProvider) as ZetaReceiverMock;
  }
}

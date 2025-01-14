/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  Callable,
  CallableInterface,
} from "../../IGatewayEVM.sol/Callable";

const _abi = [
  {
    type: "function",
    name: "onCall",
    inputs: [
      {
        name: "context",
        type: "tuple",
        internalType: "struct MessageContext",
        components: [
          {
            name: "sender",
            type: "address",
            internalType: "address",
          },
        ],
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "payable",
  },
] as const;

export class Callable__factory {
  static readonly abi = _abi;
  static createInterface(): CallableInterface {
    return new Interface(_abi) as CallableInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): Callable {
    return new Contract(address, _abi, runner) as unknown as Callable;
  }
}

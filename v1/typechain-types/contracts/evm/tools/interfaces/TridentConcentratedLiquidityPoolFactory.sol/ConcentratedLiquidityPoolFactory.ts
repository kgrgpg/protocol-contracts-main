/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../../common";

export interface ConcentratedLiquidityPoolFactoryInterface
  extends utils.Interface {
  functions: {
    "getPools(address,address,uint256,uint256)": FunctionFragment;
  };

  getFunction(nameOrSignatureOrTopic: "getPools"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "getPools",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;

  decodeFunctionResult(functionFragment: "getPools", data: BytesLike): Result;

  events: {};
}

export interface ConcentratedLiquidityPoolFactory extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ConcentratedLiquidityPoolFactoryInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    getPools(
      token0: PromiseOrValue<string>,
      token1: PromiseOrValue<string>,
      startIndex: PromiseOrValue<BigNumberish>,
      count: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[string[]] & { pairPools: string[] }>;
  };

  getPools(
    token0: PromiseOrValue<string>,
    token1: PromiseOrValue<string>,
    startIndex: PromiseOrValue<BigNumberish>,
    count: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<string[]>;

  callStatic: {
    getPools(
      token0: PromiseOrValue<string>,
      token1: PromiseOrValue<string>,
      startIndex: PromiseOrValue<BigNumberish>,
      count: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<string[]>;
  };

  filters: {};

  estimateGas: {
    getPools(
      token0: PromiseOrValue<string>,
      token1: PromiseOrValue<string>,
      startIndex: PromiseOrValue<BigNumberish>,
      count: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    getPools(
      token0: PromiseOrValue<string>,
      token1: PromiseOrValue<string>,
      startIndex: PromiseOrValue<BigNumberish>,
      count: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}

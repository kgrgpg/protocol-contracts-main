import addresses from "../data/addresses.json";

export declare type ZetaProtocolAddress =
  | "connector"
  | "erc20Custody"
  | "immutableCreate2Factory"
  | "tss"
  | "tssUpdater"
  | "zetaToken"
  | "zetaTokenConsumerUniV2"
  | "zetaTokenConsumerUniV3";

export const zetaProtocolAddress: ZetaProtocolAddress[] = [
  "connector",
  "erc20Custody",
  "immutableCreate2Factory",
  "tss",
  "tssUpdater",
  "zetaToken",
  "zetaTokenConsumerUniV2",
  "zetaTokenConsumerUniV3",
];
export const isZetaProtocolAddress = (str: string): str is ZetaProtocolAddress =>
  zetaProtocolAddress.includes(str as ZetaProtocolAddress);

export declare type ZetaZEVMAddress =
  | "fungibleModule"
  | "systemContract"
  | "uniswapv2Factory"
  | "uniswapv2Router02"
  | "zrc20";

export declare type ZetaProtocolTestNetwork =
  | "amoy_testnet"
  | "bsc_testnet"
  | "btc_testnet"
  | "sepolia_testnet"
  | "base_testnet"
  | "zeta_testnet";

export const zetaProtocolTestNetworks: ZetaProtocolTestNetwork[] = [
  "amoy_testnet",
  "bsc_testnet",
  "btc_testnet",
  "sepolia_testnet",
  "base_testnet",
  "zeta_testnet",
];

export declare type NonZetaAddress =
  | "uniswapV2Factory"
  | "uniswapV2Router02"
  | "uniswapV3Factory"
  | "uniswapV3Router"
  | "weth9";

export const nonZetaAddress: NonZetaAddress[] = [
  "uniswapV2Factory",
  "uniswapV2Router02",
  "uniswapV3Router",
  "uniswapV3Factory",
  "weth9",
];

export declare type ZetaProtocolMainNetwork =
  | "bsc_mainnet"
  | "btc_mainnet"
  | "eth_mainnet"
  | "polygon_mainnet"
  | "base_mainnet"
  | "zeta_mainnet";

export const zetaProtocolMainNetworks: ZetaProtocolMainNetwork[] = [
  "bsc_mainnet",
  "btc_mainnet",
  "eth_mainnet",
  "polygon_mainnet",
  "base_mainnet",
  "zeta_mainnet",
];

export declare type ZetaProtocolNetwork = ZetaProtocolMainNetwork | ZetaProtocolTestNetwork;
export const zetaProtocolNetworks: ZetaProtocolNetwork[] = [...zetaProtocolTestNetworks, ...zetaProtocolMainNetworks];

export declare type ZetaProtocolEnvironment = "mainnet" | "testnet";

export const isProtocolNetworkName = (str: string): str is ZetaProtocolNetwork =>
  zetaProtocolNetworks.includes(str as ZetaProtocolNetwork);

export const isTestnetNetwork = (network: string): boolean => {
  return zetaProtocolTestNetworks.includes(network as ZetaProtocolTestNetwork);
};

export const isMainnetNetwork = (network: string): boolean => {
  return zetaProtocolMainNetworks.includes(network as ZetaProtocolMainNetwork);
};

export const getZRC20Address = (network: ZetaProtocolNetwork): string => {
  return (addresses["zevm"] as any)[network]["zrc20"];
};

export const getNonZetaAddress = (address: NonZetaAddress, network: ZetaProtocolNetwork): string => {
  return (addresses["non_zeta"] as any)[network][address];
};

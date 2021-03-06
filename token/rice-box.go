package token

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "mainnet-tokens.json",
		FileModTime: time.Unix(1604945427, 0),

		Content: string("[\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x476c5E26a75bd202a9683ffD34359C0CC15be0fF/logo.png\",\n        \"mintAddress\": \"SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt\",\n        \"tokenName\": \"Serum\",\n        \"tokenSymbol\": \"SRM\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x476c5E26a75bd202a9683ffD34359C0CC15be0fF/logo.png\",\n        \"mintAddress\": \"MSRMcoVyrFxnSgo5uXwone5SKcGhT1KEJMFEkMEWf9L\",\n        \"tokenName\": \"MegaSerum\",\n        \"tokenSymbol\": \"MSRM\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/bitcoin/info/logo.png\",\n        \"mintAddress\": \"9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E\",\n        \"tokenName\": \"Wrapped Bitcoin\",\n        \"tokenSymbol\": \"BTC\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2/logo.png\",\n        \"mintAddress\": \"2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk\",\n        \"tokenName\": \"Wrapped Ethereum\",\n        \"tokenSymbol\": \"ETH\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/f3ffd0b9ae2165336279ce2f8db1981a55ce30f8/blockchains/ethereum/assets/0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9/logo.png\",\n        \"mintAddress\": \"AGFEad2et2ZJif9jaGpdMixQqvW5i81aBdvKe7PHNfz3\",\n        \"tokenName\": \"Wrapped FTT\",\n        \"tokenSymbol\": \"FTT\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.png\",\n        \"mintAddress\": \"3JSf5tPeuscJGtaCp5giEiDhv51gQ4v3zWg8DGgyLfAB\",\n        \"tokenName\": \"Wrapped YFI\",\n        \"tokenSymbol\": \"YFI\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x514910771AF9Ca656af840dff83E8264EcF986CA/logo.png\",\n        \"mintAddress\": \"CWE8jPTUYhdCTZYWPTe1o5DFqfdjzWKc9WKz6rSjQUdG\",\n        \"tokenName\": \"Wrapped Chainlink\",\n        \"tokenSymbol\": \"LINK\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ripple/info/logo.png\",\n        \"mintAddress\": \"Ga2AXHpfAF6mv2ekZwcsJFqu7wB4NV331qNH7fW9Nst8\",\n        \"tokenName\": \"Wrapped XRP\",\n        \"tokenSymbol\": \"XRP\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/f3ffd0b9ae2165336279ce2f8db1981a55ce30f8/blockchains/ethereum/assets/0xdAC17F958D2ee523a2206206994597C13D831ec7/logo.png\",\n        \"mintAddress\": \"BQcdHdAQW1hczDbBi9hiegXAR7A98Q9jx3X3iBBBDiq4\",\n        \"tokenName\": \"Wrapped USDT\",\n        \"tokenSymbol\": \"USDT\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/f3ffd0b9ae2165336279ce2f8db1981a55ce30f8/blockchains/ethereum/assets/0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48/logo.png\",\n        \"mintAddress\": \"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v\",\n        \"tokenName\": \"USD Coin\",\n        \"tokenSymbol\": \"USDC\"\n    },\n    {\n        \"deprecated\": true,\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/f3ffd0b9ae2165336279ce2f8db1981a55ce30f8/blockchains/ethereum/assets/0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48/logo.png\",\n        \"mintAddress\": \"BXXkv6z8ykpG1yuvUDPgh732wzVHB69RnB9YgSYh3itW\",\n        \"tokenName\": \"Wrapped USDC\",\n        \"tokenSymbol\": \"WUSDC\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x6B3595068778DD592e39A122f4f5a5cF09C90fE2/logo.png\",\n        \"mintAddress\": \"AR1Mtgh7zAtxuxGd2XPovXPVjcSdY3i4rQYisNadjfKy\",\n        \"tokenName\": \"Wrapped SUSHI\",\n        \"tokenSymbol\": \"SUSHI\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/6996a371cd02f516506a8f092eeb29888501447c/blockchains/nuls/assets/NULSd6HgyZkiqLnBzTaeSQfx1TNg2cqbzq51h/logo.png\",\n        \"mintAddress\": \"CsZ5LZkDS7h9TDKjrbL7VAwQZ9nsRu8vJLhRYfmGaN8K\",\n        \"tokenName\": \"Wrapped ALEPH\",\n        \"tokenSymbol\": \"ALEPH\"\n    },\n    {\n        \"icon\": \"https://github.com/trustwallet/assets/raw/b0ab88654fe64848da80d982945e4db06e197d4f/blockchains/ethereum/assets/0x8CE9137d39326AD0cD6491fb5CC0CbA0e089b6A9/logo.png\",\n        \"mintAddress\": \"SF3oTvfWzEP3DTwGSvUXRrGTvr75pdZNnBLAH9bzMuX\",\n        \"tokenName\": \"Wrapped SXP\",\n        \"tokenSymbol\": \"SXP\"\n    },\n    {\n        \"mintAddress\": \"BtZQfWqDGbk9Wf2rXEiWyQBdBY1etnUUn6zEphvVS7yN\",\n        \"tokenName\": \"Wrapped HGET\",\n        \"tokenSymbol\": \"HGET\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/4c82c2a409f18a4dd96a504f967a55a8fe47026d/blockchains/smartchain/assets/0xd4CB328A82bDf5f03eB737f37Fa6B370aef3e888/logo.png\",\n        \"mintAddress\": \"5Fu5UUgbjpUvdBveb3a1JTNirL8rXtiYeSMWvKjtUNQv\",\n        \"tokenName\": \"Wrapped CREAM\",\n        \"tokenSymbol\": \"CREAM\"\n    },\n    {\n        \"mintAddress\": \"873KLxCbz7s9Kc4ZzgYRtNmhfkQrhfyWGZJBmyCbC3ei\",\n        \"tokenName\": \"Wrapped UBXT\",\n        \"tokenSymbol\": \"UBXT\"\n    },\n    {\n        \"mintAddress\": \"HqB7uswoVg4suaQiDP3wjxob1G5WdZ144zhdStwMCq7e\",\n        \"tokenName\": \"Wrapped HNT\",\n        \"tokenSymbol\": \"HNT\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/6e375e4e5fb0ffe09ed001bae1ef8ca1d6c86034/blockchains/ethereum/assets/0xf8C3527CC04340b208C854E985240c02F7B7793f/logo.png\",\n        \"mintAddress\": \"9S4t2NEAiJVMvPdRYKVrfJpBafPBLtvbvyS3DecojQHw\",\n        \"tokenName\": \"Wrapped FRONT\",\n        \"tokenSymbol\": \"FRONT\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/878dcab0fab90e6593bcb9b7d941be4915f287dc/blockchains/ethereum/assets/0xb2734a4Cec32C81FDE26B0024Ad3ceB8C9b34037/logo.png\",\n        \"mintAddress\": \"6WNVCuxCGJzNjmMZoKyhZJwvJ5tYpsLyAtagzYASqBoF\",\n        \"tokenName\": \"Wrapped AKRO\",\n        \"tokenSymbol\": \"AKRO\"\n    },\n    {\n        \"mintAddress\": \"DJafV9qemGp7mLMEn5wrfqaFwxsbLgUsGVS16zKRk9kc\",\n        \"tokenName\": \"Wrapped HXRO\",\n        \"tokenSymbol\": \"HXRO\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/08d734b5e6ec95227dc50efef3a9cdfea4c398a1/blockchains/ethereum/assets/0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984/logo.png\",\n        \"mintAddress\": \"DEhAasscXF4kEGxFgJ3bq4PpVGp5wyUxMRvn6TzGVHaw\",\n        \"tokenName\": \"Wrapped UNI\",\n        \"tokenSymbol\": \"UNI\"\n    },\n    {\n        \"mintAddress\": \"GeDS162t9yGJuLEHPWXXGrb1zwkzinCgRwnT8vHYjKza\",\n        \"tokenName\": \"Wrapped MATH\",\n        \"tokenSymbol\": \"MATH\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/08d734b5e6ec95227dc50efef3a9cdfea4c398a1/blockchains/tomochain/info/logo.png\",\n        \"mintAddress\": \"GXMvfY2jpQctDqZ9RoU3oWPhufKiCcFEfchvYumtX7jd\",\n        \"tokenName\": \"Wrapped TOMO\",\n        \"tokenSymbol\": \"TOMO\"\n    },\n    {\n        \"icon\": \"https://raw.githubusercontent.com/trustwallet/assets/2d2491130e6beda208ba4fc6df028a82a0106ab6/blockchains/ethereum/assets/0xB1f66997A5760428D3a87D68b90BfE0aE64121cC/logo.png\",\n        \"mintAddress\": \"EqWCKXfs3x47uVosDpTRgFniThL9Y8iCztJaapxbEaVX\",\n        \"tokenName\": \"Wrapped LUA\",\n        \"tokenSymbol\": \"LUA\"\n    }\n]\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1604945427, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "mainnet-tokens.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`mints-data`, &embedded.EmbeddedBox{
		Name: `mints-data`,
		Time: time.Unix(1604945427, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"mainnet-tokens.json": file2,
		},
	})
}

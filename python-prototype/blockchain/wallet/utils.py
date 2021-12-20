from bip_utils import (Bip32Secp256k1, Bip39Languages, Bip39MnemonicGenerator,
                       Bip39MnemonicValidator, Bip39SeedGenerator,
                       Bip39WordsNum)

LANGUAGE = Bip39Languages.ENGLISH
WORDS_NUM = Bip39WordsNum.WORDS_NUM_12


class WalletUtils():

    # TODO: unit test
    @staticmethod
    def read_key_from_file(file_path: str) -> str:
        with open(file_path, 'r') as f:
            return f.read()

    @staticmethod
    def generate_mnemonic() -> str:
        mnemonic = Bip39MnemonicGenerator(LANGUAGE).FromWordsNumber(WORDS_NUM)
        return str(mnemonic)

    @staticmethod
    def is_valid_mnemonic(mnemonic: str) -> bool:
        return Bip39MnemonicValidator().IsValid(mnemonic)

    @staticmethod
    def generate_seed_bytes(mnemonic: str) -> bytes:
        return Bip39SeedGenerator(mnemonic).Generate()

    @staticmethod
    def generate_seed_bytes_string(mnemonic: str) -> str:
        seed_bytes = WalletUtils.generate_seed_bytes(mnemonic)
        return seed_bytes.hex()

    @staticmethod
    def construct_private_key(seed_bytes: bytes):
        return Bip32Secp256k1.FromSeed(seed_bytes).PrivateKey().ToExtended()

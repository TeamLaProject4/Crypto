from typing import List, Union

from bip_utils import Bip39MnemonicGenerator, Bip39WordsNum, Bip39Languages, Bip39SeedGenerator, Bip32Secp256k1, Bip39MnemonicValidator
from bip_utils.utils.mnemonic.mnemonic import Mnemonic


LANGUAGE = Bip39Languages.ENGLISH
WORDS_NUM = Bip39WordsNum.WORDS_NUM_12


class WalletUtils():

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

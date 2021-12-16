from typing import List

from bip_utils import Bip39MnemonicGenerator, Bip39WordsNum, Bip39Languages, Bip39SeedGenerator
from bip_utils.utils.mnemonic.mnemonic import Mnemonic


LANGUAGE = Bip39Languages.ENGLISH
WORDS_NUM = Bip39WordsNum.WORDS_NUM_12


class WalletUtils():

    @staticmethod
    def generate_mnemonic() -> Mnemonic:
        mnemonic = Bip39MnemonicGenerator(LANGUAGE).FromWordsNumber(WORDS_NUM)
        return mnemonic

    @staticmethod
    def generate_mnemonic_from_words(words: List[str]) -> Mnemonic:
        return Mnemonic(words)

    @staticmethod
    def get_mnemonic_words(mnemonic: Mnemonic) -> List[str]:
        return mnemonic.m_mnemonic_list

    @staticmethod
    def generate_seed_bytes(mnemonic: Mnemonic) -> bytes:
        return Bip39SeedGenerator(mnemonic).Generate()

    @staticmethod
    def generate_seed_bytes_string(mnemonic: Mnemonic) -> str:
        seed_bytes = WalletUtils.generate_seed_bytes(mnemonic)
        return seed_bytes.hex()

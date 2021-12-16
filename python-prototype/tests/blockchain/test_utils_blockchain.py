import pytest
from blockchain import BlockchainUtils

# because json_dumped, string that will be hashed is "super_cool_string" (including quotes)
TEST_STRING = 'super_cool_string'
TEST_STRING_HASH = '523d1898d54774efe136fa958634aca750d06096958ad5b59e1f9e830c57b0e1'


def test_when_string_hashed_then_HASH_MATCHES():
    hashed_string = BlockchainUtils.hash(TEST_STRING).hexdigest()

    assert hashed_string == TEST_STRING_HASH

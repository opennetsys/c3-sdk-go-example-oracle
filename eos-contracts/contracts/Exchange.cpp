#include <eosiolib/eosio.hpp>
#include <eosiolib/print.hpp>
#include <eosiolib/crypto.h>

using namespace eosio;

template<typename CharT>
static std::string to_hex(const CharT* d, uint32_t s) {
  std::string r;
  const char* to_hex="0123456789abcdef";
  uint8_t* c = (uint8_t*)d;
  for( uint32_t i = 0; i < s; ++i ) {
    (r += to_hex[(c[i] >> 4)]) += to_hex[(c[i] & 0x0f)];
  }
  return r;
}

std::string hex_to_string(const std::string& input) {
  static const char* const lut = "0123456789abcdef";
  size_t len = input.length();
  if (len & 1) abort();
  std::string output;
  output.reserve(len / 2);
  for (size_t i = 0; i < len; i += 2) {
    char a = input[i];
    const char* p = std::lower_bound(lut, lut + 16, a);
    if (*p != a) abort();
    char b = input[i + 1];
    const char* q = std::lower_bound(lut, lut + 16, b);
    if (*q != b) abort();
    output.push_back(((p - lut) << 4) | (q - lut));
  }
  return output;
}

class exchange: public eosio::contract {
  public:
      using contract::contract;

      /// @abi table checkpoint i64
      struct checkpoint {
        uint64_t id; // primary key
        std::string root; // block hash

        uint64_t primary_key() const { return id; }
        uint64_t by_checkpoint_id() const { return id; }

        EOSLIB_SERIALIZE(checkpoint, (id)(root));
      };

      typedef multi_index<N(checkpoint), checkpoint, indexed_by<N(byroot), const_mem_fun<checkpoint, uint64_t, &checkpoint::by_checkpoint_id>>> checkpoints_table;

      ///@abi action
      void placeorder(std::string root) {
        require_auth(_self);

        checkpoints_table _checkpoints(_self, _self);

        bool exists = false;
        for (auto iter = _checkpoints.begin(); iter != _checkpoints.end(); iter++) {
          if ((*iter).root == root) {
            exists = true;
          }
        }

        if (!exists) {
          _checkpoints.emplace(_self, [&](auto &row) {
              row.id = _checkpoints.available_primary_key();
              row.root = root;
              });
        } else {
          abort();
        }
      }

      void chkpointroot(std::string root) {
        require_auth(_self);

        checkpoints_table _checkpoints(_self, _self);

        bool exists = false;
        for (auto iter = _checkpoints.begin(); iter != _checkpoints.end(); iter++) {
          if ((*iter).root == root) {
            exists = true;
          }
        }

        if (!exists) {
          _checkpoints.emplace(_self, [&](auto &row) {
              row.id = _checkpoints.available_primary_key();
              row.root = root;
              });
        } else {
          abort();
        }
      }

      void getchkpoints() {
        checkpoints_table _checkpoints(_self, _self);

        for (auto iter = _checkpoints.begin(); iter != _checkpoints.end(); iter++) {
          print("id: ", (*iter).id);
          print("root: ", (*iter).root);
        }
      }
  }
};

EOSIO_ABI( exchange, (placeorder) )

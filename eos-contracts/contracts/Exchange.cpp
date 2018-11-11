#include <string>
#include <eosiolib/eosio.hpp>
#include <eosiolib/time.hpp>
#include <eosiolib/crypto.h>
#include <eosiolib/asset.hpp>
#include <eosiolib/print.h>

using namespace eosio;
using std::string;

// @abi table request i64
struct order
{

  uint64_t id; // primary key
  name sender;
  uint64_t price;
  uint64_t amount;
  uint64_t value;

  uint64_t primary_key() const { return id; }

  EOSLIB_SERIALIZE(order, (id)(sender)(price)(amount)(value))
};

// @abi table request i64
struct deposit
{

  name sender;
  uint64_t value;

  uint64_t primary_key() const { return sender; }

  EOSLIB_SERIALIZE(deposit, (value))
};

typedef multi_index<N(order), order> orders_table;
typedef multi_index<N(deposit), deposit> deposits_table;

class exchange: public eosio::contract
{
  public:
    using contract::contract;

    orders_table orders_t;
    deposits_table deposits_t;

    exchange(account_name s) : contract(s), orders_t(_self, _self), deposits_t(_self, _self) {}

    void placeorder(uint64_t price, uint64_t amount, uint64_t value)
    {
      require_auth(_self);

      auto itr = deposits_t.find(_self);
      if (itr != deposits_t.end())
      {
        deposits_t.modify(itr, _self, [&](deposit &r) {
            r.value = r.value + value;
            });
      } else {
        deposits_t.emplace(_self, [&](deposit &r) {
            r.value = value;
            });
      }

      orders_t.emplace(_self, [&](order &r) {
          r.price = price;
          r.amount = amount;
          r.value = value;
          });

    }

    void withdraw(account_name receiver, asset value)
    {
      require_auth(_self);

      account_name from = _self;

      action(
          permission_level{ from, N(active) },
          N(eosio.token), N(transfer),
          std::make_tuple(from, receiver, value, std::string(""))
          ).send();

    }
};

EOSIO_ABI(exchange, (placeorder)(withdraw))
  ~

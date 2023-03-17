create
extension pg_trgm;
/* trigram extension for faster text-search */

drop table if exists issue_NFT;
create table issue_NFT
(
    block_height bigint    not null,
    timestamp    TIMESTAMP NOT NULL DEFAULT now(),
    tx_hash      text      not null,
    class_id     text      not null,
    name         text      not null,
    sender       text      not null,
    primary key (tx_hash)
);

drop table if exists mint_NFT;
create table mint_NFT
(
    block_height bigint    not null,
    timestamp    TIMESTAMP NOT NULL DEFAULT now(),
    tx_hash      text      not null,
    class_id     text      not null,
    nft_id       text      not null,
    recipient    text      not null,
    owner        text      not null,
    token_uri    text      not null,
    primary key (tx_hash)
);

drop table if exists transfer_NFT;
create table transfer_NFT
(
    block_height  bigint    not null,
    timestamp     TIMESTAMP NOT NULL DEFAULT now(),
    tx_hash       text      not null,
    class_id      text      not null,
    base_class_id text      not null,
    nft_id        text      not null,
    sender        text      not null,
    receiver      text      not null,
    primary key (tx_hash, base_class_id, nft_id)
);
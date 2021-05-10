# nohup ./ZapMiner --config=local_cfgs/config1/config.json mine >> logs/1.log &
# nohup ./ZapMiner --config=local_cfgs/config2/config.json mine >> logs/2.log &
# nohup ./ZapMiner --config=local_cfgs/config3/config.json mine >> logs/3.log &
# nohup ./ZapMiner --config=local_cfgs/config4/config.json mine >> logs/4.log &
# nohup ./ZapMiner --config=local_cfgs/config5/config.json mine >> logs/5.log &

./zap-miner approve 10000 0xCD8a1C3ba11CF5ECfa6267617243239504a98d90
./zap-miner approve 10000 0xb7278a61aa25c888815afc32ad3cc52ff24fe575

nohup ./zap-miner --config=local_cfgs/config1.json mine -r > logs/1.log &
nohup ./zap-miner --config=local_cfgs/config2.json mine -r > logs/2.log &
nohup ./zap-miner --config=local_cfgs/config3.json mine -r > logs/3.log &
nohup ./zap-miner --config=local_cfgs/config4.json mine -r > logs/4.log &
nohup ./zap-miner --config=local_cfgs/config5.json mine -r > logs/5.log &
# /bin/bash

zip_tar=''
datetime=$(date +%Y%m%d_%H%M%S)
if grep -q 'enable_bank=true' ./run.sh; then
    zip_tar="res_preempt_${datetime}.zip"
else
    zip_tar="res_caravan_${datetime}.zip"
fi

echo 'ziping ' $zip_tar
zip -r results/$zip_tar log

cp ./results/$zip_tar /disc/home/dy/caddy/site
echo 'copyed ' $zip_tar 'to download site'

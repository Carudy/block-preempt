zip_tar=''

if grep -q 'enable_bank=true' ./run.sh; then
    zip_tar=res_preempt.zip
else
    zip_tar=res_caravan.zip
fi

echo 'ziping ' $zip_tar

rm ./$zip_tar
zip -r $zip_tar log

cp ./$zip_tar /disc/home/dy/caddy/site
echo 'copyed ' $zip_tar 'to site'

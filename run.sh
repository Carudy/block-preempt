#!/bin/bash

./clean.sh
# Configuration variables - edit these as needed
enable_bank=true
num_shards=2
num_nodes=16
malicious_num=1
inject_speed=500

if [ "$enable_bank" = "true" ]; then
    echo "BANK enabled."
    BANK_FLAG="-b"
else
    echo "Caravan mode."
    BANK_FLAG=""
fi

test_file="20W.csv"

echo "=== Starting Blockchain Shard Network ==="
echo "Shards: $num_shards"
echo "Nodes per shard: $num_nodes"
echo "Malicious nodes per shard: $malicious_num"
echo "Test file: $test_file"
echo ""


# Create log directory
log_dir="./process_logs"
mkdir -p "$log_dir"
echo "Logs will be saved to: $log_dir/"
echo ""

# Array to store PIDs
node_pids=()

echo "Starting nodes..."
for ((s=0; s<num_shards; s++)); do
    shard_id="S$s"
    for ((n=0; n<num_nodes; n++)); do
        node_id="N$n"
        log_file="$log_dir/${shard_id}_${node_id}.log"

        echo "Starting $shard_id $node_id..."

        # Start node with exact command format from requirements
        ./block_caravan_exe -i $inject_speed $BANK_FLAG -S $num_shards -N $num_nodes -f $malicious_num -s $shard_id -n $node_id -t $test_file > "$log_file" 2>&1 &

        pid=$!
        node_pids+=($pid)

        # Brief pause between node starts
        sleep 0.2

        echo "  Started (PID: $pid, Log: $log_file)"
    done
done

echo ""
echo "All nodes started!"
echo ""

# Start client/supervisor
echo "Starting client/supervisor..."
client_log="$log_dir/client.log"

# Start client with -c flag as requested
./block_caravan_exe -i $inject_speed $BANK_FLAG -c -S $num_shards -N $num_nodes -t $test_file  > "$client_log" 2>&1 &

client_pid=$!
node_pids+=($client_pid)

sleep 1
echo "Client started (PID: $client_pid, Log: $client_log)"
echo ""

echo "=== Network is running ==="
echo "Total processes: ${#node_pids[@]}"
echo "All output is being logged to: $log_dir/"
echo ""
echo "To view logs:"
echo "  tail -f $log_dir/*.log"

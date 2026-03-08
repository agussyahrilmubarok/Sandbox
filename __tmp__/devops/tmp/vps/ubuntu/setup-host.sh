#!/bin/bash

# Exit on error
set -e

# 1. Setting the Timezone
echo "Setting timezone to Asia/Jakarta..."
timedatectl set-timezone Asia/Jakarta

# 2. Setting the Hostname
HOSTNAME="${1:-$(hostname)}"  # Default to current hostname if not provided
echo "Setting hostname to ${HOSTNAME}..."
hostnamectl set-hostname "${HOSTNAME}"

# 3. Installing Required Packages
echo "Updating package lists..."
apt update
echo "Installing required packages..."
apt install -y ca-certificates curl gnupg lsb-release chrony rsync net-tools iotop sysstat python3-psycopg2

# 4. Forcing NTP Sync
echo "Forcing NTP sync..."
chronyc -a makestep

# 5. Disabling Swap Partition
echo "Disabling swap partition..."
swapoff -a

# 6. Disabling Swap Partition Permanently
echo "Disabling swap partition permanently..."
# Remove swap entries from /etc/fstab
grep -v '^.*swap\s' /etc/fstab | sudo tee /etc/fstab

echo "Setup complete."

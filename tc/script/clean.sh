#!/bin/bash

cat << EOF >> /usr/lib/systemd/system/shutdown-clean.service
[Unit]
Description=close services before reboot and shutdown
DefaultDependencies=no
Before=shutdown.target reboot.target halt.target
#Before=network.target iscsi.service iscsid.service shutdown.target reboot.target halt.target
# This works because it is installed in the target and will be executed before the 
# target state is entered
# Also consider kexec.target

[Service]
Type=oneshot
ExecStart=~/ebpf/tc/script/commit.sh

[Install]
WantedBy=halt.target reboot.target shutdown.target

EOF

systemctl daemon-reload
systemctl enable shutdown-clean.service

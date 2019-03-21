#!/usr/bin/env python3
import sys
import boto3
from datetime import datetime
import logging


LOGGER = logging.getLogger()
LOGGER.setLevel(logging.INFO)
MAX_DAYS = 14


def main():
    ec2 = boto3.resource('ec2', region_name='eu-west-1')
    client = boto3.client('ec2')
    snaps = ec2.snapshots.filter(
        Filters=[
            {'Name': 'tag:custom', 'Values': ["build"]}
        ]
    )
    if list(snaps):
        for snap in snaps:
            snap_date = snap.start_time.replace(tzinfo=None)
            current_date = datetime.now()
            delta = current_date - snap_date
            if delta.days > MAX_DAYS:
                LOGGER.info("deleting snapshot %s is %d days old",
                            snap.snapshot_id,
                            delta.days
                            )
                client.delete_snapshot(SnapshotId=snap.snapshot_id,
                                       DryRun=True
                                       )


if __name__ == '__main__':
    sys.exit(main())

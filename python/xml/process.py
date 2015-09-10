#!/usr/bin/env python
import xml.etree.ElementTree as et

parsin_data = """
?xml version="1.0"?>
<datacenter>
        <server tag="done">
                <id>101</id>
                <memory>32G</memory>
                <compute>Intel Xeon</compute>
                <network>1G</network>
        </server>
        <server tag="pending">
                <id>102</id>
                <memory>128G</memory>
                <compute>2 Intel Xeon</compute>
                <network>10G</network>
        </server>
</datacenter>
"""


def run():
    tree = et.parse("servers.xml")
    root = tree.getroot()
    print(root.tag)
    for server in root.findall('server'):
        print("server")
        id_server = server.find('id').text
        compute = server.find('compute').text
        network = server.find('network').text
        print("\t{0}".format(id_server))
        print("\t{0}".format(compute))
        print("\t{0}".format(network))
    print(root.tag)


if __name__ == '__main__':
    run()

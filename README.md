# building
make

# Testing
make test

# example  
./bin/bloxtool-go host get hostname-here

# Usage  
  Usage:  
    bloxtool record:host get <hostname> <view>  
    bloxtool record:host create <hostname> <ipv4addrs> <view> [--mac=<mac>] [--configure-for-dhcp=<true>]  
    bloxtool record:host delete <hostname> <view>  
    bloxtool record:cname get <alias> <view>  
    bloxtool record:cname create <alias> <cname> <view>  
    bloxtool record:cname update <alias> <cname> <view>  
    bloxtool record:cname delete <alias> <view>  

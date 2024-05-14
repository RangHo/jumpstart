# Jumpstart

**Jumpstart** is a collection of small [Ignition](https://coreos.github.io/ignition/) configuration files and utilities for [Fedora CoreOS](https://fedoraproject.org/coreos/) that are accessible via HTTP(S).

## Available Endpoints

### PXE Netboot

These endpoints provide an easy, semi-frozen URLs to the latest CoreOS images.
It can be used to provision a bare-metal nodes, and is especially useful to reprovision a single-node CoreOS installation.

| Endpoint                                  | Description                                                                                                                                                                                                      |
|-------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `/pxe/kernel?stream=<stream>&arch=<arch>` | Redirects to the latest PXE kernel image for the specified _stream_ and _arch_ pair. `<stream>` can be one of `stable`, `testing`, or `next`; `<arch>` can be one of `x86_64`, `aarch64`, `s390x`, or `ppc64le`. |
| `/pxe/rootfs?stream=<stream>&arch=<arch>` | Redirects to the latest PXE rootfs image for the specified _stream_ and _arch_ pair. `<stream>` can be one of `stable`, `testing`, or `next`; `<arch>` can be one of `x86_64`, `aarch64`, `s390x`, or `ppc64le`. |
| `/pxe/initramfs?stream=<stream>&arch=<arch>` | Redirects to the latest PXE initramfs image for the specified _stream_ and _arch_ pair. `<stream>` can be one of `stable`, `testing`, or `next`; `<arch>` can be one of `x86_64`, `aarch64`, `s390x`, or `ppc64le`. |

# Jumpstart

**Jumpstart** is a collection of small [Ignition](https://coreos.github.io/ignition/) configuration files and utilities for [Fedora CoreOS](https://fedoraproject.org/coreos/) that are accessible via HTTP(S).

## Available Endpoints

### Image Shortcuts

These endpoints provide an easy, semi-frozen URLs to the latest CoreOS images.
It can be used to provision a bare-metal nodes, and is especially useful to reprovision a single-node CoreOS installation.
Each image endpoints requires two parameters: `stream` and `arch`.

| Endpoint                                          | Description                                                                             |
|---------------------------------------------------|-----------------------------------------------------------------------------------------|
| `/raw?stream=<stream>&arch=<arch>`                | Redirects to the latest raw `.tar.xz` image for the specified _stream_ and _arch_ pair. |
| `/iso?stream=<stream>&arch=<arch>`                | Redirects to the latest ISO image for the specified _stream_ and _arch_ pair.           |
| `/pxe/kernel?stream=<stream>&architecture=<arch>` | Redirects to the latest PXE kernel image for the specified _stream_ and _arch_ pair.    |
| `/pxe/rootfs?stream=<stream>&arch=<arch>`         | Redirects to the latest PXE rootfs image for the specified _stream_ and _arch_ pair.    |
| `/pxe/initramfs?stream=<stream>&arch=<arch>`      | Redirects to the latest PXE initramfs image for the specified _stream_ and _arch_ pair. |

`stream` parameter specifies the release stream.
Replace the `<stream>` placeholder with one of:

- `stable` for the most reliable version of CoreOS;
- `testing` for the next stable release candidate; or
- `next` for the "bleeding edge" release.

`arch` parameter specifies the target CPU architecture.
Replace the `<arch>` placeholder with one of:

- `x86_64` for most computers with Intel or AMD processors;
- `aarch64` for ARM-powered devices such as Raspberry Pi;
- `s390x` for IBM Cloud and zSystems; or
- `ppc64le` for IBM PowerPC systems.

### User Management

These endpoints provide a dynamic user definition that allows user provisioning without manual copying of values.
All user-related endpoints reside in the `/passwd` subpath.

| Endpoint                              | Description                                                                                                                                              |
|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| `/passwd/from-github?user=<username>` | Creates the `core` user with the same SSH keys that the given GitHub user uses to authenticate on GitHub. Replace `<username>` with the GitHub username. |

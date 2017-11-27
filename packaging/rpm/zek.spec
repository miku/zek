Summary:    Generate a Go struct from an XML document.
Name:       zek
Version:    0.1.0
Release:    0
License:    GPL
BuildArch:  x86_64
BuildRoot:  %{_tmppath}/%{name}-build
Group:      System/Base
Vendor:     Leipzig University Library, https://www.ub.uni-leipzig.de
URL:        https://github.com/miku/zek

%description

Generate a Go struct from an XML document.

%prep

%build

%pre

%install

mkdir -p $RPM_BUILD_ROOT/usr/local/sbin
install -m 755 zek $RPM_BUILD_ROOT/usr/local/sbin

%post

%clean
rm -rf $RPM_BUILD_ROOT
rm -rf %{_tmppath}/%{name}
rm -rf %{_topdir}/BUILD/%{name}

%files
%defattr(-,root,root)

/usr/local/sbin/zek

%changelog

* Mon Nov 27 2017 Martin Czygan
- 0.1.0 initial release

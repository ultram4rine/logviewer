Name:       logviewer
Version:    0.1.1
Release:    1%{?dist}
Summary:    LogViewer web-application
License:    MIT
URL:        http://stingr.net
Source0:	%{name}-%{version}.tar.gz
 
BuildRequires:	golang
BuildRequires:  systemd
Requires(pre): shadow-utils
%{?systemd_requires}
 
%description
Web-application to view logs from dhcp and switches
 
%pre
getent group %{name} >/dev/null || groupadd -r %{name}
getent passwd %{name} >/dev/null || \
    useradd -r -g %{name} -d %{_sharedstatedir}/%{name} -s /sbin/nologin \
    -c "SGU Hosting Controller" %{name}
exit 0
 
%post
%systemd_post %{name}.service %{name}.socket
 
%preun
%systemd_preun %{name}.service %{name}.socket
 
%postun
%systemd_postun_with_restart %{name}.service %{name}.socket
 
%prep
%setup -q
 
%build
mkdir -p goapp/src/git.sgu.ru/ultramarine goapp/bin
ln -s ${PWD} goapp/src/git.sgu.ru/ultramarine/%{name}
export GO111MODULE=off
export GOPATH=${PWD}/goapp
%gobuild -o goapp/bin/%{name} git.sgu.ru/ultramarine/%{name}
 
%install
 
%{__install} -d $RPM_BUILD_ROOT%{_bindir}
%{__install} -v -D -t $RPM_BUILD_ROOT%{_bindir} goapp/bin/%{name}
%{__install} -d $RPM_BUILD_ROOT%{_unitdir}
%{__install} -v -D -t $RPM_BUILD_ROOT%{_unitdir} %{name}.service
%{__install} -v -D -t $RPM_BUILD_ROOT%{_unitdir} %{name}.socket
%{__install} -d -m 0755 %{buildroot}%{_sysconfdir}/%{name}
%{__install} -d $RPM_BUILD_ROOT%{_sysconfdir}/sysconfig
%{__install} -m 644 -T %{name}.sysconfig %{buildroot}%{_sysconfdir}/sysconfig/%{name}
%{__install} -d -m 0755 %{buildroot}%{_sharedstatedir}/%{name}
 
%files
%{_bindir}/%{name}
%{_unitdir}/%{name}.service
%{_unitdir}/%{name}.socket
%config(noreplace) %{_sysconfdir}/%{name}
%config(noreplace) %{_sysconfdir}/sysconfig/%{name}
%dir %attr(-,%{name},%{name}) %{_sharedstatedir}/%{name}
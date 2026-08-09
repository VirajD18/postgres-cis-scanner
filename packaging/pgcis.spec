Name:           pgcis
Version:        1.0.0
Release:        1%{?dist}
Summary:        PostgreSQL CIS Security Scanner

License:        Proprietary
BuildArch:      x86_64

%description
PostgreSQL CIS Benchmark Security Scanner.

%prep

%build

%install
mkdir -p %{buildroot}/usr/bin
mkdir -p %{buildroot}/etc/pgcis
mkdir -p %{buildroot}/usr/share/pgcis/benchmark
mkdir -p %{buildroot}/var/lib/pgcis/reports
mkdir -p %{buildroot}/etc/pgcis/templates

install -m 0755 %{_sourcedir}/pgcis \
    %{buildroot}/usr/bin/pgcis

install -m 0600 %{_sourcedir}/servers.json \
    %{buildroot}/etc/pgcis/servers.json

install -m 0644 %{_sourcedir}/README.txt \
    %{buildroot}/etc/pgcis/README.txt

install -m 0644 %{_sourcedir}/iaas.json \
    %{buildroot}/etc/pgcis/templates/iaas.json

install -m 0644 %{_sourcedir}/pass.json \
    %{buildroot}/etc/pgcis/templates/pass.json

cp -r %{_sourcedir}/benchmark/* \
    %{buildroot}/usr/share/pgcis/benchmark/

%files
/usr/bin/pgcis
%config(noreplace) /etc/pgcis/servers.json
%config(noreplace) /etc/pgcis/README.txt
%config(noreplace) /etc/pgcis/templates/iaas.json
%config(noreplace) /etc/pgcis/templates/pass.json
/usr/share/pgcis/benchmark
/var/lib/pgcis/reports

%changelog
* Sun Aug 09 2026 Viraj Deshmukh - 1.0.0-1
- Initial PostgreSQL CIS Scanner release

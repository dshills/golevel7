my @files=<*.tsv>;
print << 'END';
package commons;

var FieldNames = map[string][]string{
END

foreach my $file (@files) {
	if ($file=~/^(...)\.tsv/i) {
		print "\t\"".uc($1)."\": []string{\n";
		print "\t\t\"".uc($1)." Record\",\n";
		open(FH,"<$file");
		while (<FH>) {
			my $line=$_;
			chomp $line;
			my @flds=split(/\t/,$line);
			my $fldDesc=$flds[0];
			if (my $matches=$fldDesc=~/^[0-9]*\-(.*)$/) {
				print "\t\t\"$1\",\n";
			}
		}
		print "\t},\n";
	}
}

print "}\n";


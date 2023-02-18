Parse and filter place visits from Google Location History.

## Usage

You'll need to download the location history data from [Google Takeout](https://takeout.google.com).

Just point `lh` at the downloaded zip file and it will print location info out:

    lh -i takeout-20230210T211622Z-001.zip

You can also filter the data by date:

    lh -i takeout-20230210T211622Z-001.zip -f 2011-10-22
    
or by partial date:
    
    lh -i takeout-20230210T211622Z-001.zip -f 2011-10
    

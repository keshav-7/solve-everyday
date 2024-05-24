<?php

    $conn = mysqli_connect('localhost', 'root', 'localdb123', 'elms');

    if (!$conn) {
        die("Connection failed: " . mysqli_connect_error());
    }
    die("Connection success!");    
    function getTatInfo(int $labId = 0)
    {
        // $dataArr = (new MerchantCosting())->getTatInfo($labId);
        $dataArr = getTatInfoSQL($labId);
        print_r($dataArr);die;
    }


    function getTatInfoSQL(int $labId): ?array
    {
        $sql = 'SELECT
            mc.deal_type,
            mc.deal_type_id,
            mc.merchant_id,
            mc.cutoff_time,
            mc.same_day_delivery,
            mc.next_day_delivery,
            mc.sunday,
            mc.monday,
            mc.tuesday,
            mc.wednesday,
            mc.thursday,
            mc.friday,
            mc.saturday,
            dc.city_id,
            czm.zone_id,
            mc.travel_time,
            dc.operational_tat
        FROM
            merchant_costing AS mc
        LEFT JOIN 
            sample_collection_zone AS scz ON scz.labId = mc.merchant_id
        LEFT JOIN
            city_zone_mapping AS czm ON czm.zone_id = scz.zone_id
        LEFT JOIN 
            deal_city AS dc ON dc.city_id = czm.city_id
        WHERE 
            mc.merchant_id = ' . $labId . '
            AND mc.cutoff_time IS NOT NULL
            AND mc.same_day_delivery IS NOT NULL
            AND mc.next_day_delivery IS NOT NULL
            AND mc.sunday + mc.monday + mc.tuesday + mc.wednesday + mc.thursday + mc.friday + mc.saturday > 0
            AND scz.isactive = 1
            AND czm.status = 1
            AND dc.city_status = 1 LIMIT 100';
        return $this->db->query($sql)->getResultArray();
    }
?>

